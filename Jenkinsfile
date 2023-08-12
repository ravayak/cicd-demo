node {
    //  la fonction tool permet de définir des configurations d'outils dans Jenkins
    //  elle permet de déclarer et d'utiliser des outils externes tels que les SDK, les compilateurs, les interpréteurs, etc.
    def root = tool type: 'go', name: '1.20'

    withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
        stage("clone"){
             checkout scm 
        }
        stage("build"){
            sh "go build -v -o app ."
            sh "git tag -l build-jenkins"
            sh "git tag -a -f -m 'Jenkins build' build-jenkins"
            sh "git --version"
            sh 'docker build . -t ynno/go-cicd-demo'
        }
        stage('deliver') {
            steps {
                withCredentials([usernamePassword(
                    credentialsId: 'GITHUB_TOKEN_ID',
                    passwordVariable: 'TOKEN',
                    usernameVariable: 'USER')]) {
                    sh 'git push https://${USER}:${TOKEN}@github.com/ravayak/cicd-demo.git build-jenkins'
            }

                withCredentials([usernamePassword(credentialsId: 'dockerhub', passwordVariable: 'dockerhubPassword', usernameVariable: 'dockerhubUser')]) {
                    sh "docker login -u ${env.dockerhubUser} -p ${env.dockerhubPassword}"
                    sh 'docker push ynno/go-cicd-demo'
                }
            }
        }
    }
}