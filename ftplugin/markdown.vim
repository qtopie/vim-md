" VimMdPreview
" VimMdImagePaste
" VimMdImageClean

" Workaround for https://github.com/vim/vim/issues/4530
if exists("g:vimmdpluginloaded")
  finish
endif
let g:vimmdpluginloaded=1

augroup vimmd
augroup END

let s:channel = ""
"let s:timer = ""
"let s:plugindir = expand(expand("<sfile>:p:h:h"))

function s:define(channel, msg)
  try
  catch
    let l:resp[2][0] = 'Caught ' . string(v:exception) . ' in ' . v:throwpoint
  endtry
  call ch_sendexpr(a:channel, l:resp)
endfunction

function s:doShutdown() 
  call ch_close(s:channel)
endfunction

let opts = {"in_mode": "json", "out_mode": "json", "err_mode": "json", "callback": function("s:define"), "timeout": 30000} 
let targetdir = "/home/qtopierw/.vim/plugged/vim-md/"
let start = targetdir."vim-md"
" TODO remove me for debug purpose only
echo start
let job = job_start(start, opts)
let s:channel = job_getchannel(job)

" TODO remove me for debug purpose only
echo "job started"

au VimLeave * call s:doShutdown()