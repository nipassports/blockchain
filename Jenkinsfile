pipeline {
  agent any
  stages {
    stage('Stop') {
      parallel {
        stage('Stop Fabric') {
          steps {
            sh 'docker rm -f $(docker ps | grep "hyperledger/fabric" | cut -d" " -f1) | exit 0'
            sh 'docker rm -f $(docker ps | grep "dev-peer" | cut -d" " -f1) | exit 0'
          }
        }
        stage('Clean Fabric') {
          steps {
            sh 'docker rmi -f $(docker images | grep "hyperledger/fabric" | cut -d" " -f1) | exit 0'
            sh 'docker rmi -f $(docker images | grep passport | cut -d" " -f1) | exit 0'
            sh 'docker rmi -f $(docker images | grep visa | cut -d" " -f1) | exit 0'
            sh 'docker rmi -f $(docker images | grep "hyperledger/fabric" | cut -d" " -f1) | exit 0'
            sh 'docker rmi -f $(docker images | grep "dev-peer" | cut -d" " -f1) | exit 0'
          }
        }
        stage('Stop back-end') {
          steps {
            sh 'docker rm -f -v nip-back | exit 0'
          }
        }
      }
    }
    stage('Build') {
      steps {
        sh './create_bin.sh'
      }
    }
    stage('Start Fabric') {
      steps {
        sh 'cd nip && ./startFabric.sh'
        sh 'cd nip && ./actionsPostStart.sh'
      }
    }
    stage('Restart back') {
      steps {
        sh '''docker run -d --name nip-back  -v /var/lib/jenkins/workspace/blockchain_master/nip-network/:/app/nip-network -v /var/lib/jenkins/workspace/blockchain_master/nip/javascript/wallet:/app/nip/javascript/wallet --net=host nip/back
'''
      }
    }
  }
}
