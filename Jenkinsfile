node {
    def root = tool type: 'go', name: '1.20'

    withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
        stage("clone"){
             checkout scm 
             sh "git checkout main"
        }
        stage("build"){
            sh "go build -v -o app ."
            sh "git tag -l build-jenkins"
            sh "git tag -a -f -m 'Jenkins build' build-jenkins"
            sh "git --version"
            withCredentials([usernamePassword(
                credentialsId: 'GITHUB_TOKEN_ID',
                passwordVariable: 'TOKEN',
                usernameVariable: 'USER')]) {
            sh "git push https://antonny.kihm@protonmail.com:${TOKEN}@github.com/ravayak/cicd-demo.git build-jenkins"
            }
        }
    }
}