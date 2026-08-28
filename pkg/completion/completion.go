package completion

const BashCompletion = `
_gleedos_completion() {
    local cur prev opts
    COMPREPLY=()
    cur="${COMP_WORDS[COMP_CWORD]}"
    prev="${COMP_WORDS[COMP_CWORD-1]}"
    opts="--best --audio --turbo --threads --limit-rate --cookies --subs --batch --watch serve history doctor"

    if [[ ${cur} == -* ]] ; then
        COMPREPLY=( $(compgen -W "${opts}" -- ${cur}) )
        return 0
    fi
}
complete -F _gleedos_completion gleedos
`

const ZshCompletion = `
#compdef gleedos
_gleedos() {
    _arguments -s         '--best[Best available quality]'         '--audio[Extract audio as MP3]'         '--turbo[Boost concurrency]'         '--threads[Parallel threads]:threads:'         '--limit-rate[Rate limit]:rate:'         '--batch[Batch file]:file:_files'         '--watch[Clipboard watcher]'         '*:urls:_urls'
}
`
