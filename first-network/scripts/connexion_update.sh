#!/bin/bash




KEYORG1=$(ls ../crypto-config/peerOrganizations/org1.nip.ddns.net/users/Admin@org1.nip.ddns.net/msp/keystore/)
KEYORG2=$(ls ../crypto-config/peerOrganizations/org2.nip.ddns.net/users/Admin@org2.nip.ddns.net/msp/keystore/)
KEYORG3=$(ls ../crypto-config/peerOrganizations/org3.nip.ddns.net/users/Admin@org3.nip.ddns.net/msp/keystore/)

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
        \"nipchannel\": {
            \"orderers\": [
                \"orderer.nip.ddns.net\"
            ],
            \"peers\": {
                \"peer0.org1.nip.ddns.net\": {
                    \"endorsingPeer\": true,
                    \"chaincodeQuery\": true,
                    \"eventSource\": true
                },
                \"peer1.org1.nip.ddns.net\": {
                    \"endorsingPeer\": true,
                    \"chaincodeQuery\": true,
                    \"eventSource\": true
                },
                \"peer0.org2.nip.ddns.net\": {
                    \"endorsingPeer\": true,
                    \"chaincodeQuery\": true,
                    \"eventSource\": true
                },
                \"peer1.org2.nip.ddns.net\": {
                    \"endorsingPeer\": true,
                    \"chaincodeQuery\": true,
                    \"eventSource\": true
                },
                \"peer0.org3.nip.ddns.net\": {
                    \"endorsingPeer\": true,
                    \"chaincodeQuery\": true,
                    \"eventSource\": true
                },
                \"peer1.org3.nip.ddns.net\": {
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
                \"peer0.org1.nip.ddns.net\",
                \"peer1.org1.nip.ddns.net\"
            ],
            \"certificateAuthorities\": [
                \"ca.org1.nip.ddns.net\"
            ],
            \"adminPrivateKey\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org1.nip.ddns.net/users/Admin@org1.nip.ddns.net/msp/keystore/$KEYORG1\"
            },
            \"signedCert\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org1.nip.ddns.net/users/Admin@org1.nip.ddns.net/msp/signcerts/Admin@org1.nip.ddns.net-cert.pem\"
            }
        },
        \"Org2\": {
            \"mspid\": \"Org2MSP\",
            \"peers\": [
                \"peer0.org2.nip.ddns.net\",
                \"peer1.org2.nip.ddns.net\"
            ],
            \"certificateAuthorities\": [
                \"ca.org2.nip.ddns.net\"
            ],
            \"adminPrivateKey\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org2.nip.ddns.net/users/Admin@org2.nip.ddns.net/msp/keystore/$KEYORG2\"
            },
            \"signedCert\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org2.nip.ddns.net/users/Admin@org2.nip.ddns.net/msp/signcerts/Admin@org2.nip.ddns.net-cert.pem\"
            }
        },
        \"Org3\": {
            \"mspid\": \"Org3MSP\",
            \"peers\": [
                \"peer0.org3.nip.ddns.net\",
                \"peer1.org3.nip.ddns.net\"
            ],
            \"certificateAuthorities\": [
                \"ca.org3.nip.ddns.net\"
            ],
            \"adminPrivateKey\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org3.nip.ddns.net/users/Admin@org3.nip.ddns.net/msp/keystore/$KEYORG3\"
            },
            \"signedCert\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org3.nip.ddns.net/users/Admin@org3.nip.ddns.net/msp/signcerts/Admin@org3.nip.ddns.net-cert.pem\"
            }
        }
    },
    \"orderers\": {
        \"orderer.nip.ddns.net\": {
            \"url\": \"grpcs://localhost:7050\",
            \"grpcOptions\": {
                \"ssl-target-name-override\": \"orderer.nip.ddns.net\"
            },
            \"tlsCACerts\": {
                \"path\": \"../../first-network/crypto-config/ordererOrganizations/nip.ddns.net/tlsca/tlsca.nip.ddns.net-cert.pem\"
            }
        }
    },
    \"peers\": {
        \"peer0.org1.nip.ddns.net\": {
            \"url\": \"grpcs://localhost:7051\",
            \"eventUrl\": \"grpcs://localhost:7053\",
            \"grpcOptions\": {
                \"ssl-target-name-override\": \"peer0.org1.nip.ddns.net\"
            },
            \"tlsCACerts\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org1.nip.ddns.net/tlsca/tlsca.org1.nip.ddns.net-cert.pem\"
            }
        },
        \"peer1.org1.nip.ddns.net\": {
            \"url\": \"grpcs://localhost:8051\",
            \"eventUrl\": \"grpcs://localhost:8053\",
            \"grpcOptions\": {
                \"ssl-target-name-override\": \"peer1.org1.nip.ddns.net\"
            },
            \"tlsCACerts\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org1.nip.ddns.net/tlsca/tlsca.org1.nip.ddns.net-cert.pem\"
            }
        },
        \"peer0.org2.nip.ddns.net\": {
            \"url\": \"grpcs://localhost:9051\",
            \"eventUrl\": \"grpcs://localhost:9053\",
            \"grpcOptions\": {
                \"ssl-target-name-override\": \"peer0.org2.nip.ddns.net\"
            },
            \"tlsCACerts\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org2.nip.ddns.net/tlsca/tlsca.org2.nip.ddns.net-cert.pem\"
            }
        },
        \"peer1.org2.nip.ddns.net\": {
            \"url\": \"grpcs://localhost:10051\",
            \"eventUrl\": \"grpcs://localhost:10053\",
            \"grpcOptions\": {
                \"ssl-target-name-override\": \"peer1.org2.nip.ddns.net\"
            },
            \"tlsCACerts\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org2.nip.ddns.net/tlsca/tlsca.org2.nip.ddns.net-cert.pem\"
            }
        },
        \"peer0.org3.nip.ddns.net\": {
            \"url\": \"grpcs://localhost:11051\",
            \"eventUrl\": \"grpcs://localhost:11053\",
            \"grpcOptions\": {
                \"ssl-target-name-override\": \"peer0.org3.nip.ddns.net\"
            },
            \"tlsCACerts\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org3.nip.ddns.net/tlsca/tlsca.org3.nip.ddns.net-cert.pem\"
            }
        },
        \"peer1.org3.nip.ddns.net\": {
            \"url\": \"grpcs://localhost:12051\",
            \"eventUrl\": \"grpcs://localhost:12053\",
            \"grpcOptions\": {
                \"ssl-target-name-override\": \"peer1.org3.nip.ddns.net\"
            },
            \"tlsCACerts\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org3.nip.ddns.net/tlsca/tlsca.org3.nip.ddns.net-cert.pem\"
            }
        }
    },
    \"certificateAuthorities\": {
        \"ca.org1.nip.ddns.net\": {
            \"url\": \"https://localhost:7054\",
            \"caName\": \"ca-org1\",
            \"httpOptions\": {
                \"verify\": false
            },
            \"tlsCACerts\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org1.nip.ddns.net/ca/ca.org1.nip.ddns.net-cert.pem\"
            }

        },
        \"ca.org2.nip.ddns.net\": {
            \"url\": \"https://localhost:8054\",
            \"caName\": \"ca-org2\",
            \"httpOptions\": {
                \"verify\": false
            },
            \"tlsCACerts\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org2.nip.ddns.net/ca/ca.org2.nip.ddns.net-cert.pem\"
            }
        },
        \"ca.org3.nip.ddns.net\": {
            \"url\": \"https://localhost:9054\",
            \"caName\": \"ca-org3\",
            \"httpOptions\": {
                \"verify\": false
            },
            \"tlsCACerts\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org3.nip.ddns.net/ca/ca.org3.nip.ddns.net-cert.pem\"
            }

        }
    }
}

" > connection1.json


echo  "{
    \"name\": \"byfn-network\",
    \"x-type\": \"hlfv1\",
    \"version\": \"1.0.0\",
    \"client\": {
        \"organization\": \"Org2\",
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
        \"nipchannel\": {
            \"orderers\": [
                \"orderer.nip.ddns.net\"
            ],
            \"peers\": {
                \"peer0.org1.nip.ddns.net\": {
                    \"endorsingPeer\": true,
                    \"chaincodeQuery\": true,
                    \"eventSource\": true
                },
                \"peer1.org1.nip.ddns.net\": {
                    \"endorsingPeer\": true,
                    \"chaincodeQuery\": true,
                    \"eventSource\": true
                },
                \"peer0.org2.nip.ddns.net\": {
                    \"endorsingPeer\": true,
                    \"chaincodeQuery\": true,
                    \"eventSource\": true
                },
                \"peer1.org2.nip.ddns.net\": {
                    \"endorsingPeer\": true,
                    \"chaincodeQuery\": true,
                    \"eventSource\": true
                },
                \"peer0.org3.nip.ddns.net\": {
                    \"endorsingPeer\": true,
                    \"chaincodeQuery\": true,
                    \"eventSource\": true
                },
                \"peer1.org3.nip.ddns.net\": {
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
                \"peer0.org1.nip.ddns.net\",
                \"peer1.org1.nip.ddns.net\"
            ],
            \"certificateAuthorities\": [
                \"ca.org1.nip.ddns.net\"
            ],
            \"adminPrivateKey\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org1.nip.ddns.net/users/Admin@org1.nip.ddns.net/msp/keystore/$KEYORG1\"
            },
            \"signedCert\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org1.nip.ddns.net/users/Admin@org1.nip.ddns.net/msp/signcerts/Admin@org1.nip.ddns.net-cert.pem\"
            }
        },
        \"Org2\": {
            \"mspid\": \"Org2MSP\",
            \"peers\": [
                \"peer0.org2.nip.ddns.net\",
                \"peer1.org2.nip.ddns.net\"
            ],
            \"certificateAuthorities\": [
                \"ca.org2.nip.ddns.net\"
            ],
            \"adminPrivateKey\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org2.nip.ddns.net/users/Admin@org2.nip.ddns.net/msp/keystore/$KEYORG2\"
            },
            \"signedCert\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org2.nip.ddns.net/users/Admin@org2.nip.ddns.net/msp/signcerts/Admin@org2.nip.ddns.net-cert.pem\"
            }
        },
        \"Org3\": {
            \"mspid\": \"Org3MSP\",
            \"peers\": [
                \"peer0.org3.nip.ddns.net\",
                \"peer1.org3.nip.ddns.net\"
            ],
            \"certificateAuthorities\": [
                \"ca.org3.nip.ddns.net\"
            ],
            \"adminPrivateKey\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org3.nip.ddns.net/users/Admin@org3.nip.ddns.net/msp/keystore/$KEYORG3\"
            },
            \"signedCert\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org3.nip.ddns.net/users/Admin@org3.nip.ddns.net/msp/signcerts/Admin@org3.nip.ddns.net-cert.pem\"
            }
        }
    },
    \"orderers\": {
        \"orderer.nip.ddns.net\": {
            \"url\": \"grpcs://localhost:7050\",
            \"grpcOptions\": {
                \"ssl-target-name-override\": \"orderer.nip.ddns.net\"
            },
            \"tlsCACerts\": {
                \"path\": \"../../first-network/crypto-config/ordererOrganizations/nip.ddns.net/tlsca/tlsca.nip.ddns.net-cert.pem\"
            }
        }
    },
    \"peers\": {
        \"peer0.org1.nip.ddns.net\": {
            \"url\": \"grpcs://localhost:7051\",
            \"eventUrl\": \"grpcs://localhost:7053\",
            \"grpcOptions\": {
                \"ssl-target-name-override\": \"peer0.org1.nip.ddns.net\"
            },
            \"tlsCACerts\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org1.nip.ddns.net/tlsca/tlsca.org1.nip.ddns.net-cert.pem\"
            }
        },
        \"peer1.org1.nip.ddns.net\": {
            \"url\": \"grpcs://localhost:8051\",
            \"eventUrl\": \"grpcs://localhost:8053\",
            \"grpcOptions\": {
                \"ssl-target-name-override\": \"peer1.org1.nip.ddns.net\"
            },
            \"tlsCACerts\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org1.nip.ddns.net/tlsca/tlsca.org1.nip.ddns.net-cert.pem\"
            }
        },
        \"peer0.org2.nip.ddns.net\": {
            \"url\": \"grpcs://localhost:9051\",
            \"eventUrl\": \"grpcs://localhost:9053\",
            \"grpcOptions\": {
                \"ssl-target-name-override\": \"peer0.org2.nip.ddns.net\"
            },
            \"tlsCACerts\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org2.nip.ddns.net/tlsca/tlsca.org2.nip.ddns.net-cert.pem\"
            }
        },
        \"peer1.org2.nip.ddns.net\": {
            \"url\": \"grpcs://localhost:10051\",
            \"eventUrl\": \"grpcs://localhost:10053\",
            \"grpcOptions\": {
                \"ssl-target-name-override\": \"peer1.org2.nip.ddns.net\"
            },
            \"tlsCACerts\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org2.nip.ddns.net/tlsca/tlsca.org2.nip.ddns.net-cert.pem\"
            }
        },
        \"peer0.org3.nip.ddns.net\": {
            \"url\": \"grpcs://localhost:11051\",
            \"eventUrl\": \"grpcs://localhost:11053\",
            \"grpcOptions\": {
                \"ssl-target-name-override\": \"peer0.org3.nip.ddns.net\"
            },
            \"tlsCACerts\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org3.nip.ddns.net/tlsca/tlsca.org3.nip.ddns.net-cert.pem\"
            }
        },
        \"peer1.org3.nip.ddns.net\": {
            \"url\": \"grpcs://localhost:12051\",
            \"eventUrl\": \"grpcs://localhost:12053\",
            \"grpcOptions\": {
                \"ssl-target-name-override\": \"peer1.org3.nip.ddns.net\"
            },
            \"tlsCACerts\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org3.nip.ddns.net/tlsca/tlsca.org3.nip.ddns.net-cert.pem\"
            }
        }
    },
    \"certificateAuthorities\": {
        \"ca.org1.nip.ddns.net\": {
            \"url\": \"https://localhost:7054\",
            \"caName\": \"ca-org1\",
            \"httpOptions\": {
                \"verify\": false
            },
            \"tlsCACerts\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org1.nip.ddns.net/ca/ca.org1.nip.ddns.net-cert.pem\"
            }

        },
        \"ca.org2.nip.ddns.net\": {
            \"url\": \"https://localhost:8054\",
            \"caName\": \"ca-org2\",
            \"httpOptions\": {
                \"verify\": false
            },
            \"tlsCACerts\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org2.nip.ddns.net/ca/ca.org2.nip.ddns.net-cert.pem\"
            }
        },
        \"ca.org3.nip.ddns.net\": {
            \"url\": \"https://localhost:9054\",
            \"caName\": \"ca-org3\",
            \"httpOptions\": {
                \"verify\": false
            },
            \"tlsCACerts\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org3.nip.ddns.net/ca/ca.org3.nip.ddns.net-cert.pem\"
            }

        }
    }
}

" > connection2.json

echo  "{
    \"name\": \"byfn-network\",
    \"x-type\": \"hlfv1\",
    \"version\": \"1.0.0\",
    \"client\": {
        \"organization\": \"Org3\",
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
        \"nipchannel\": {
            \"orderers\": [
                \"orderer.nip.ddns.net\"
            ],
            \"peers\": {
                \"peer0.org1.nip.ddns.net\": {
                    \"endorsingPeer\": true,
                    \"chaincodeQuery\": true,
                    \"eventSource\": true
                },
                \"peer1.org1.nip.ddns.net\": {
                    \"endorsingPeer\": true,
                    \"chaincodeQuery\": true,
                    \"eventSource\": true
                },
                \"peer0.org2.nip.ddns.net\": {
                    \"endorsingPeer\": true,
                    \"chaincodeQuery\": true,
                    \"eventSource\": true
                },
                \"peer1.org2.nip.ddns.net\": {
                    \"endorsingPeer\": true,
                    \"chaincodeQuery\": true,
                    \"eventSource\": true
                },
                \"peer0.org3.nip.ddns.net\": {
                    \"endorsingPeer\": true,
                    \"chaincodeQuery\": true,
                    \"eventSource\": true
                },
                \"peer1.org3.nip.ddns.net\": {
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
                \"peer0.org1.nip.ddns.net\",
                \"peer1.org1.nip.ddns.net\"
            ],
            \"certificateAuthorities\": [
                \"ca.org1.nip.ddns.net\"
            ],
            \"adminPrivateKey\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org1.nip.ddns.net/users/Admin@org1.nip.ddns.net/msp/keystore/$KEYORG1\"
            },
            \"signedCert\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org1.nip.ddns.net/users/Admin@org1.nip.ddns.net/msp/signcerts/Admin@org1.nip.ddns.net-cert.pem\"
            }
        },
        \"Org2\": {
            \"mspid\": \"Org2MSP\",
            \"peers\": [
                \"peer0.org2.nip.ddns.net\",
                \"peer1.org2.nip.ddns.net\"
            ],
            \"certificateAuthorities\": [
                \"ca.org2.nip.ddns.net\"
            ],
            \"adminPrivateKey\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org2.nip.ddns.net/users/Admin@org2.nip.ddns.net/msp/keystore/$KEYORG2\"
            },
            \"signedCert\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org2.nip.ddns.net/users/Admin@org2.nip.ddns.net/msp/signcerts/Admin@org2.nip.ddns.net-cert.pem\"
            }
        },
        \"Org3\": {
            \"mspid\": \"Org3MSP\",
            \"peers\": [
                \"peer0.org3.nip.ddns.net\",
                \"peer1.org3.nip.ddns.net\"
            ],
            \"certificateAuthorities\": [
                \"ca.org3.nip.ddns.net\"
            ],
            \"adminPrivateKey\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org3.nip.ddns.net/users/Admin@org3.nip.ddns.net/msp/keystore/$KEYORG3\"
            },
            \"signedCert\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org3.nip.ddns.net/users/Admin@org3.nip.ddns.net/msp/signcerts/Admin@org3.nip.ddns.net-cert.pem\"
            }
        }
    },
    \"orderers\": {
        \"orderer.nip.ddns.net\": {
            \"url\": \"grpcs://localhost:7050\",
            \"grpcOptions\": {
                \"ssl-target-name-override\": \"orderer.nip.ddns.net\"
            },
            \"tlsCACerts\": {
                \"path\": \"../../first-network/crypto-config/ordererOrganizations/nip.ddns.net/tlsca/tlsca.nip.ddns.net-cert.pem\"
            }
        }
    },
    \"peers\": {
        \"peer0.org1.nip.ddns.net\": {
            \"url\": \"grpcs://localhost:7051\",
            \"eventUrl\": \"grpcs://localhost:7053\",
            \"grpcOptions\": {
                \"ssl-target-name-override\": \"peer0.org1.nip.ddns.net\"
            },
            \"tlsCACerts\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org1.nip.ddns.net/tlsca/tlsca.org1.nip.ddns.net-cert.pem\"
            }
        },
        \"peer1.org1.nip.ddns.net\": {
            \"url\": \"grpcs://localhost:8051\",
            \"eventUrl\": \"grpcs://localhost:8053\",
            \"grpcOptions\": {
                \"ssl-target-name-override\": \"peer1.org1.nip.ddns.net\"
            },
            \"tlsCACerts\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org1.nip.ddns.net/tlsca/tlsca.org1.nip.ddns.net-cert.pem\"
            }
        },
        \"peer0.org2.nip.ddns.net\": {
            \"url\": \"grpcs://localhost:9051\",
            \"eventUrl\": \"grpcs://localhost:9053\",
            \"grpcOptions\": {
                \"ssl-target-name-override\": \"peer0.org2.nip.ddns.net\"
            },
            \"tlsCACerts\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org2.nip.ddns.net/tlsca/tlsca.org2.nip.ddns.net-cert.pem\"
            }
        },
        \"peer1.org2.nip.ddns.net\": {
            \"url\": \"grpcs://localhost:10051\",
            \"eventUrl\": \"grpcs://localhost:10053\",
            \"grpcOptions\": {
                \"ssl-target-name-override\": \"peer1.org2.nip.ddns.net\"
            },
            \"tlsCACerts\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org2.nip.ddns.net/tlsca/tlsca.org2.nip.ddns.net-cert.pem\"
            }
        },
        \"peer0.org3.nip.ddns.net\": {
            \"url\": \"grpcs://localhost:11051\",
            \"eventUrl\": \"grpcs://localhost:11053\",
            \"grpcOptions\": {
                \"ssl-target-name-override\": \"peer0.org3.nip.ddns.net\"
            },
            \"tlsCACerts\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org3.nip.ddns.net/tlsca/tlsca.org3.nip.ddns.net-cert.pem\"
            }
        },
        \"peer1.org3.nip.ddns.net\": {
            \"url\": \"grpcs://localhost:12051\",
            \"eventUrl\": \"grpcs://localhost:12053\",
            \"grpcOptions\": {
                \"ssl-target-name-override\": \"peer1.org3.nip.ddns.net\"
            },
            \"tlsCACerts\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org3.nip.ddns.net/tlsca/tlsca.org3.nip.ddns.net-cert.pem\"
            }
        }
    },
    \"certificateAuthorities\": {
        \"ca.org1.nip.ddns.net\": {
            \"url\": \"https://localhost:7054\",
            \"caName\": \"ca-org1\",
            \"httpOptions\": {
                \"verify\": false
            },
            \"tlsCACerts\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org1.nip.ddns.net/ca/ca.org1.nip.ddns.net-cert.pem\"
            }

        },
        \"ca.org2.nip.ddns.net\": {
            \"url\": \"https://localhost:8054\",
            \"caName\": \"ca-org2\",
            \"httpOptions\": {
                \"verify\": false
            },
            \"tlsCACerts\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org2.nip.ddns.net/ca/ca.org2.nip.ddns.net-cert.pem\"
            }
        },
        \"ca.org3.nip.ddns.net\": {
            \"url\": \"https://localhost:9054\",
            \"caName\": \"ca-org3\",
            \"httpOptions\": {
                \"verify\": false
            },
            \"tlsCACerts\": {
                \"path\": \"../../first-network/crypto-config/peerOrganizations/org3.nip.ddns.net/ca/ca.org3.nip.ddns.net-cert.pem\"
            }

        }
    }
}
" > connection3.json


exit 0
