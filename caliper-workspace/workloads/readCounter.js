'use strict';

const { WorkloadModuleBase } = require('@hyperledger/caliper-core');

class MyWorkload extends WorkloadModuleBase {
    constructor() {
        super();
    }

    async initializeWorkloadModule(workerIndex, totalWorkers, roundIndex, roundArguments, sutAdapter, sutContext) {
        await super.initializeWorkloadModule(workerIndex, totalWorkers, roundIndex, roundArguments, sutAdapter, sutContext);


    const request = {
        contractId: this.roundArguments.contractId,
        contractFunction: 'IncrementCounter',
        invokerIdentity: 'User1',
        contractArguments: null,
        readOnly: false
    };

    await this.sutAdapter.sendRequests(request);
    }

    async submitTransaction() {
        const myAgs = {
            contractId: this.roundArguments.contractId,
            contractFunction: 'GetCounter',
            invokerIdentity: 'User1',
            contractArguments: null,
            readOnly: true
        }
    
        await this.sutAdapter.sendRequests(myAgs);
    }

    async cleanupWorkloadModule() {
        const request = {
            contractId: this.roundArguments.contractId,
            contractFunction: 'DerementCounter',
            invokerIdentity: 'User1',
            contractArguments: null,
            readOnly: false
        };
    
        await this.sutAdapter.sendRequests(request);
    }
}

function createWorkloadModule() {
    return new MyWorkload();
}

module.exports.createWorkloadModule = createWorkloadModule;