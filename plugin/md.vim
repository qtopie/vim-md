let s:plugindir = expand('<sfile>:p:h:h')

" These commands are available on any filetypes
command! -nargs=* -complete=customlist,s:complete VimMdInstall call s:VimMdInstallBinaries(-1, <f-args>)
command! -nargs=* -complete=customlist,s:complete VimMdUpdate  call s:VimMdInstallBinaries(1, <f-args>)

function s:VimMdInstallBinaries(updateBinaries, ...)
  let binary = ".bin/vim-md"

  silent !clear
  execute "silent !" . "cd " . s:plugindir . "; " . "go build -o" . " " . binary
  echomsg "update vim-md plugin"
endfunction