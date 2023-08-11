node {
    def root = tool type: 'go', name: '1.20'

    withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
        stage("clone"){
            git clone 'https://github.com/ravayak/cicd-demo.git'
        }
        sh 'go version'
    }
}