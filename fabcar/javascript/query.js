/*
 * SPDX-License-Identifier: Apache-2.0
 */

'use strict';

const { FileSystemWallet, Gateway } = require('fabric-network');
const fs = require('fs');
const path = require('path');

const ccpPath = path.resolve(__dirname, '..', '..', 'first-network', 'connection3.json');
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
            console.log('An identity for the user "user3" does not exist in the wallet');
            console.log('Run the registerUser.js application before retrying');
            return;
        }

        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();
        await gateway.connect(ccp, { wallet, identity: 'user3', discovery: { enabled: false } });

        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork('mychannel');

        // Get the contract from the network.
        const contract = await network.getContract('mycc');

        // Evaluate the specified transaction.
        
        const result1 = await contract.evaluateTransaction('queryAllPassports');
        console.log(`Transaction has been evaluated - querry all passeports, result is: ${result1.toString()}\n`);

        await contract.submitTransaction('createPassport', 'P', 'EN', "45ML44147", 'brad', 'davincy','10/04/1985', 'France', 'M', 'Toulouse','1.65','Préfecture de ', 'Avenue des Facultés, 33400 Talence', 'Marron', '16/02/2023','25/01/2015','France', 'Valide','Password3', 'Image');

        const result2 = await contract.evaluateTransaction('searchPassport', 'EN');
        console.log(`Transaction has been evaluated - search passport, result is: ${result2.toString()}\n`);

    } catch (error) {
        console.error(`Failed to evaluate transaction: ${error}`);
        process.exit(1);
    }
}

main();