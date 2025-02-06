package template

const GitignoreOutput = `.gitignore`

const GitignoreTemplate = `
# docker data dir
docker/*/

# log file
*.log
*.log_lock
*.log_symlink

# goland idea
.idea

# project run dir
*run/

# go work
go.work
go.work.sum

# executable files
*.exe
`
