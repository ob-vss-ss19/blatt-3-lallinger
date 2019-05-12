pipeline {
    agent none
    stages {
        stage('Build') {
            agent {
                docker { image 'obraun/vss-protoactor-jenkins' }
            }
            steps {
                sh 'go get github.com/ob-vss-ss19/blatt-3-lallinger/messages'
                sh 'go get github.com/ob-vss-ss19/blatt-3-lallinger/tree'
                sh 'go get github.com/ob-vss-ss19/blatt-3-lallinger/treecli'
                sh 'go get github.com/ob-vss-ss19/blatt-3-lallinger/treeserver'
                sh 'cd messages && make regenerate'
                sh 'cd treeservice && go build main.go'
                sh 'cd treecli && go build main.go'
            }
        }
        stage('Test') {
            agent {
                docker { image 'obraun/vss-protoactor-jenkins' }
            }
            steps {
                sh 'echo run tests...'
            }
        }
        stage('Lint') {
            agent {
                docker { image 'obraun/vss-protoactor-jenkins' }
            }   
            steps {
                sh 'golangci-lint run --deadline 20m --enable-all'
            }
        }
        stage('Build Docker Image') {
            agent any
            steps {
                sh "docker-build-and-push -b ${BRANCH_NAME} -s treeservice -f treeservice.dockerfile"
                sh "docker-build-and-push -b ${BRANCH_NAME} -s treecli -f treecli.dockerfile"
            }
        }
    }
}
