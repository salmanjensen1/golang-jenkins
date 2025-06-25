// Jenkinsfile (Declarative Pipeline)
pipeline {
   environment {
        PATH="$PATH:/usr/local/bin:/usr/local/go/bin"
    }
 
    agent any
    stages {
        stage('build') {
            steps {
                sh '''
                echo $PATH
                which docker
                docker --version
                '''
                sh 'go version'
                echo 'Hello world'
                sh '''
                    echo "Multi Line"
                    timeout 20s top
                    ls -lah
                '''
            }
        }
    }
}
