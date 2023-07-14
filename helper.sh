#!bin/bash
ws="/Users/ashutoshyadav/projects/easy-pay" # change this to your workspace path

alias be="cd $ws/BE"
alias fe="cd $ws/FE"

alias admfe="cd $ws/FE/fe-admin-panel"
alias cstfe="cd $ws/FE/fe-easy-pay"

# Run All Projects
runall() {
    # Open new terminal and run backend project
    open -a Terminal.app && sleep 2 && osascript -e 'tell application "Terminal" to do script "cd /Users/ashutoshyadav/projects/easy-pay/BE && go run main.go"'

    # Open new terminal and run frontend admin project
    open -a Terminal.app && sleep 2 && osascript -e 'tell application "Terminal" to do script "cd /Users/ashutoshyadav/projects/easy-pay/FE/fe-admin-panel && npm run dev"'

    # Open new terminal and run frontend customer project
    open -a Terminal.app && sleep 2 && osascript -e 'tell application "Terminal" to do script "cd /Users/ashutoshyadav/projects/easy-pay/FE/fe-easy-pay && npm run dev"'
}
