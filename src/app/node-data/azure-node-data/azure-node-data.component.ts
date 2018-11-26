import {Component, Input, OnChanges, OnDestroy, OnInit, SimpleChanges} from '@angular/core';
import {FormArray, FormControl, FormGroup, Validators} from '@angular/forms';
import {Subscription} from 'rxjs';
import {ApiService, DatacenterService, WizardService} from '../../core/services';
import {AddNodeService} from '../../core/services/add-node/add-node.service';
import {CloudSpec} from '../../shared/entity/ClusterEntity';
import {DataCenterEntity} from '../../shared/entity/DatacenterEntity';
import {AzureSizes} from '../../shared/entity/provider/azure/AzureSizeEntity';
import {NodeData, NodeProviderData} from '../../shared/model/NodeSpecChange';

@Component({
  selector: 'kubermatic-azure-node-data',
  templateUrl: './azure-node-data.component.html',
  styleUrls: ['./azure-node-data.component.scss'],
})

export class AzureNodeDataComponent implements OnInit, OnDestroy, OnChanges {
  @Input() cloudSpec: CloudSpec;
  @Input() nodeData: NodeData;
  @Input() projectId: string;
  @Input() clusterId: string;
  @Input() seedDCName: string;

  sizes: AzureSizes;
  azureNodeForm: FormGroup;
  tags: FormArray;
  datacenter: DataCenterEntity;
  hideOptional = true;
  private subscriptions: Subscription[] = [];

  constructor(
      private addNodeService: AddNodeService, private wizardService: WizardService, private api: ApiService,
      private dcService: DatacenterService) {}

  ngOnInit(): void {
    const tagList = new FormArray([]);
    for (const i in this.nodeData.spec.cloud.azure.tags) {
      if (this.nodeData.spec.cloud.azure.tags.hasOwnProperty(i)) {
        tagList.push(new FormGroup({
          key: new FormControl(i),
          value: new FormControl(this.nodeData.spec.cloud.azure.tags[i]),
        }));
      }
    }

    this.azureNodeForm = new FormGroup({
      size: new FormControl(this.nodeData.spec.cloud.azure.size, Validators.required),
      assignPublicIP: new FormControl(this.nodeData.spec.cloud.azure.assignPublicIP),
      tags: tagList,
    });

    this.subscriptions.push(this.azureNodeForm.valueChanges.subscribe((data) => {
      this.addNodeService.changeNodeProviderData(this.getNodeProviderData());
    }));

    this.subscriptions.push(this.wizardService.clusterSettingsFormViewChanged$.subscribe((data) => {
      this.hideOptional = data.hideOptional;
    }));

    this.getDatacenter();
    this.reloadAzureSizes();
    this.addNodeService.changeNodeProviderData(this.getNodeProviderData());
  }

  getDatacenter(): void {
    if (this.cloudSpec.dc) {
      this.subscriptions.push(this.dcService.getDataCenter(this.cloudSpec.dc).subscribe((data) => {
        this.datacenter = data;
      }));
    }
  }

  isInWizard(): boolean {
    return !this.clusterId || this.clusterId.length === 0;
  }

  reloadAzureSizes(): void {
    if (this.cloudSpec.dc) {
      if (this.isInWizard()) {
        if (this.cloudSpec.azure.clientID && this.cloudSpec.azure.clientSecret && this.cloudSpec.azure.subscriptionID &&
            this.cloudSpec.azure.tenantID) {
          this.subscriptions.push(this.api
                                      .getAzureSizesForWizard(
                                          this.cloudSpec.azure.clientID, this.cloudSpec.azure.clientSecret,
                                          this.cloudSpec.azure.subscriptionID, this.cloudSpec.azure.tenantID,
                                          this.datacenter.spec.azure.location)
                                      .subscribe((data) => {
                                        this.sizes = data;
                                        this.azureNodeForm.controls.size.setValue(this.nodeData.spec.cloud.azure.size);
                                      }));
        }
      } else {
        this.subscriptions.push(
            this.api.getAzureSizes(this.projectId, this.seedDCName, this.clusterId).subscribe((data) => {
              this.sizes = data;
              this.azureNodeForm.controls.size.setValue(this.nodeData.spec.cloud.azure.size);
            }));
      }
    } else {
      this.getDatacenter();
    }
  }

  ngOnChanges(changes: SimpleChanges): void {
    if (changes.cloudSpec && !changes.cloudSpec.firstChange) {
      if (!!!changes.cloudSpec.previousValue ||
          (changes.cloudSpec.currentValue.azure.clientID !== changes.cloudSpec.previousValue.azure.clientID) ||
          (changes.cloudSpec.currentValue.azure.clientSecret !== changes.cloudSpec.previousValue.azure.clientSecret) ||
          (changes.cloudSpec.currentValue.azure.subscriptionID !==
           changes.cloudSpec.previousValue.azure.subscriptionID) ||
          (changes.cloudSpec.currentValue.azure.tenantID !== changes.cloudSpec.previousValue.azure.tenantID)) {
        this.reloadAzureSizes();
      }
    }
  }

  ngOnDestroy(): void {
    for (const sub of this.subscriptions) {
      if (sub) {
        sub.unsubscribe();
      }
    }
  }

  getTagForm(form): any {
    return form.get('tags').controls;
  }

  addTag(): void {
    this.tags = this.azureNodeForm.get('tags') as FormArray;
    this.tags.push(new FormGroup({
      key: new FormControl(''),
      value: new FormControl(''),
    }));
  }

  deleteTag(index: number): void {
    const arrayControl = this.azureNodeForm.get('tags') as FormArray;
    arrayControl.removeAt(index);
  }

  getNodeProviderData(): NodeProviderData {
    const tagMap = {};
    for (const i in this.azureNodeForm.controls.tags.value) {
      if (this.azureNodeForm.controls.tags.value[i].key !== '' &&
          this.azureNodeForm.controls.tags.value[i].value !== '') {
        tagMap[this.azureNodeForm.controls.tags.value[i].key] = this.azureNodeForm.controls.tags.value[i].value;
      }
    }

    return {
      spec: {
        azure: {
          size: this.azureNodeForm.controls.size.value,
          assignPublicIP: this.azureNodeForm.controls.assignPublicIP.value,
          tags: tagMap,
        },
      },
      valid: this.azureNodeForm.valid,
    };
  }
}
