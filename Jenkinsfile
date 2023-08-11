node {
    def root = tool type: 'go', name: '1.20'
    - name: Configure Git 
      run: |
        echo "machine github.com login ravayak password ${GH_TOKEN}" > ~/.netrc
        git config --global url."https://github.com/".insteadOf "git://github.com/"
        git config --global advice.detachedHead false
      env:
        GH_TOKEN: ${{ secrets.GH_TOKEN }}

    withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
        stage("clone"){
             checkout scm 
             sh "git checkout main"
        }
        stage("build"){
            sh "go build -v -o app ."
            sh "git add ."
            sh "git commit -m 'built app'"
            sh "git push"
        }
    }
}