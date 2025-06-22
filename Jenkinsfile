Jenkinsfile (Declarative Pipeline)
pipeline {
    agent { docker { image 'golang:1.24.4-alpine3.22' } }
    stages {
        stage('build') {
            steps {
                sh 'go version'
                sh 'Hello world'
                sh '''
                    echo "Multi Line"
                    htop
                    ls -lah
                '''
            }
        }
        stage('retry/timeout'){
            steps{
                retry(3){
                    sh './flakey-deploy.sh'
                }
                timeout(time:3, unit:'MINUTES'){
                    sh './health-check.sh'
                }
            }
        }
    }
}
