#!/bin/bash




KEYORG1=$(ls ../crypto-config/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp/keystore/)
KEYORG2=$(ls ../crypto-config/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp/keystore/)
KEYORG3=$(ls ../org3-artifacts/crypto-config/peerOrganizations/org3.example.com/users/Admin@org3.example.com/msp/keystore/)

cd ..
echo "{
    \"name\": \"byfn-network\",
    \"x-type\": \"hlfv1\",
    \"version\": \"1.0.0\",
    \"client\": {
        \"organization\": \"Org1\",
        \"connection\": {
            \"timeout\": {
                \"peer\": {
                    \"endorser\": \"300\"
                },
                \"orderer\": \"300\"
            }
        },
    \"credentialStore\": {
        \"path\": \"./hfc-key-store\",
        \"cryptoStore\": {
        \"path\": \"./hfc-key-store\"
        }
    }
    },
    \"tlsEnable\": true,
    \"channels\": {
        \"mychannel\": {
            \"orderers\": [
                \"orderer.example.com\"
            ],
            \"peers\": {
                \"peer0.org1.example.com\": {
                    \"endorsingPeer\": true,
                    \"chaincodeQuery\": true,
                    \"eventSource\": true
                },
                \"peer1.org1.example.com\": {
                    \"endorsingPeer\": true,
                    \"chaincodeQuery\": true,
                    \"eventSource\": true
                },
                \"peer0.org2.example.com\": {
                    \"endorsingPeer\": true,
                    \"chaincodeQuery\": true,
                    \"eventSource\": true
                },
                \"peer1.org2.example.com\": {
                    \"endorsingPeer\": true,
                    \"chaincodeQuery\": true,
                    \"eventSource\": true
                }, 
                \"peer0.org3.example.com\": {
                    \"endorsingPeer\": true,
                    \"chaincodeQuery\": true,
                    \"eventSource\": true
                }, 
                \"peer1.org3.example.com\": {
                    \"endorsingPeer\": true,
                    \"chaincodeQuery\": true,
                    \"eventSource\": true
                }
            }
        }
    },
    \"organizations\": {
        \"Org1\": {
            \"mspid\": \"Org1MSP\",
            \"peers\": [
                \"peer0.org1.example.com\",
                \"peer1.org1.example.com\"
            ],
            \"certificateAuthorities\": [
                \"ca.org1.example.com\"
            ], 
            \"adminPrivateKey\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp/keystore/$KEYORG1\"
            }, 
            \"signedCert\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp/signcerts/Admin@org1.example.com-cert.pem\"
            }
        },
        \"Org2\": {
            \"mspid\": \"Org2MSP\",
            \"peers\": [
                \"peer0.org2.example.com\",
                \"peer1.org2.example.com\"
            ],
            \"certificateAuthorities\": [
                \"ca.org2.example.com\"
            ],
            \"adminPrivateKey\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp/keystore/$KEYORG2\"
            }, 
            \"signedCert\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp/signcerts/Admin@org2.example.com-cert.pem\"
            }
        }, 
        \"Org3\": {
            \"mspid\": \"Org3MSP\",
            \"peers\": [
                \"peer0.org3.example.com\",
                \"peer1.org3.example.com\"
            ],
            \"certificateAuthorities\": [
                \"ca.org3.example.com\"
            ],
            \"adminPrivateKey\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org3.example.com/users/Admin@org3.example.com/msp/keystore/$KEYORG3\"
            }, 
            \"signedCert\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org3.example.com/users/Admin@org3.example.com/msp/signcerts/Admin@org3.example.com-cert.pem\"
            }
        }
    },
    \"orderers\": {
        \"orderer.example.com\": {
            \"url\": \"grpcs://localhost:7050\",
            \"grpcOptions\": {
                \"ssl-target-name-override\": \"orderer.example.com\"
            },
            \"tlsCACerts\": {
                \"path\": \"../../first-network/crypto-config/ordererOrganizations/example.com/tlsca/tlsca.example.com-cert.pem\"
            }
        }
    },
    \"peers\": {
        \"peer0.org1.example.com\": {
            \"url\": \"grpcs://localhost:7051\",
            \"eventUrl\": \"grpcs://localhost:7053\",
            \"grpcOptions\": {
                \"ssl-target-name-override\": \"peer0.org1.example.com\"
            },
            \"tlsCACerts\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org1.example.com/tlsca/tlsca.org1.example.com-cert.pem\"
            }
        },
        \"peer1.org1.example.com\": {
            \"url\": \"grpcs://localhost:8051\",
            \"eventUrl\": \"grpcs://localhost:8053\",
            \"grpcOptions\": {
                \"ssl-target-name-override\": \"peer1.org1.example.com\"
            },
            \"tlsCACerts\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org1.example.com/tlsca/tlsca.org1.example.com-cert.pem\"
            }
        },
        \"peer0.org2.example.com\": {
            \"url\": \"grpcs://localhost:9051\",
            \"eventUrl\": \"grpcs://localhost:9053\",
            \"grpcOptions\": {
                \"ssl-target-name-override\": \"peer0.org2.example.com\"
            },
            \"tlsCACerts\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org2.example.com/tlsca/tlsca.org2.example.com-cert.pem\"
            }
        },
        \"peer1.org2.example.com\": {
            \"url\": \"grpcs://localhost:10051\",
            \"eventUrl\": \"grpcs://localhost:10053\",
            \"grpcOptions\": {
                \"ssl-target-name-override\": \"peer1.org2.example.com\"
            },
            \"tlsCACerts\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org2.example.com/tlsca/tlsca.org2.example.com-cert.pem\"
            }
        },
        \"peer0.org3.example.com\": {
            \"url\": \"grpcs://localhost:11051\",
            \"eventUrl\": \"grpcs://localhost:11053\",
            \"grpcOptions\": {
                \"ssl-target-name-override\": \"peer0.org3.example.com\"
            },
            \"tlsCACerts\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org3.example.com/tlsca/tlsca.org3.example.com-cert.pem\"
            }
        },
        \"peer1.org3.example.com\": {
            \"url\": \"grpcs://localhost:12051\",
            \"eventUrl\": \"grpcs://localhost:12053\",
            \"grpcOptions\": {
                \"ssl-target-name-override\": \"peer1.org3.example.com\"
            },
            \"tlsCACerts\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org3.example.com/tlsca/tlsca.org3.example.com-cert.pem\"
            }
        }
    },
    \"certificateAuthorities\": {
        \"ca.org1.example.com\": {
            \"url\": \"https://localhost:7054\",
            \"caName\": \"ca-org1\",
            \"httpOptions\": {
                \"verify\": false
            }, 
            \"tlsCACerts\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org1.example.com/ca/ca.org1.example.com-cert.pem\"
            }

        },
        \"ca.org2.example.com\": {
            \"url\": \"https://localhost:8054\",
            \"caName\": \"ca-org2\",
            \"httpOptions\": {
                \"verify\": false
            }, 
            \"tlsCACerts\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org2.example.com/ca/ca.org2.example.com-cert.pem\"
            }
        },
        \"ca.org3.example.com\": {
            \"url\": \"https://localhost:9054\",
            \"caName\": \"ca-org3\",
            \"httpOptions\": {
                \"verify\": false
            }, 
            \"tlsCACerts\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org3.example.com/ca/ca.org3.example.com-cert.pem\"
            }

        }
    }
}

" > connection.json



exit 0