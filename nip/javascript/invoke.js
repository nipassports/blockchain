/*
 * SPDX-License-Identifier: Apache-2.0
 */

'use strict';

const { FileSystemWallet, Gateway } = require('fabric-network');
const fs = require('fs');
const path = require('path');

const ccpPath = path.resolve(__dirname, '..', '..', 'nip-network', 'connection3.json');
const ccpJSON = fs.readFileSync(ccpPath, 'utf8');
const ccp = JSON.parse(ccpJSON);

async function main() {
    try {

        // Create a new file system based wallet for managing identities.
        const walletPath = path.join(process.cwd(), 'wallet');
        const wallet = new FileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);

        // Check to see if we've already enrolled the user.
        const userExists = await wallet.exists('user3');
        if (!userExists) {
            console.log('An identity for the user "admin" does not exist in the wallet');
            console.log('Run the registerUser.js application before retrying');
            return;
        }

        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();
        await gateway.connect(ccp, { wallet, identity: 'user3', discovery: { enabled: true } });

        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork('nipchannel');

        // Get the contract from the network.
        const contract = network.getContract('passport');
         //const contract = await network.getContract('visa');

        // Submit the specified transaction.
        //await contract.submitTransaction('changePassport' ,'PKKKKK', 'FR', "14ML52147", 'brazsad', 'davszazsincy','12/4/1995', 'France', 'M', 'Toulouse','1.65','Préfecture de ', 'Avenue des Facultés, 33400 Talence', 'Marron', '16/02/2023','25/01/2015','France', 'Valide', 'Image');
        //await contract.submitTransaction('changePassportValidity', '2');
        //await contract.submitTransaction('createVisa', 'P', '14ML521497', 'fds', 'Jean', 'Dupont', 'ddf', '16/09/1985', 'France', 'Toulouse','dfs', 'Préfecture de ', 'Avenue des Facultés, 33400 Talence', 'Marron', '16/02/2023');
        await contract.submitTransaction('changePassword','14ML66146', 'okaaaaay');


        console.log('Transaction has been submitted');

        // Disconnect from the gateway.
        await gateway.disconnect();

    } catch (error) {
        console.error(`Failed to submit transaction: ${error}`);
        process.exit(1);
    }
}

main();
