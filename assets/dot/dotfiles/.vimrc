" basic config
syntax on
filetype plugin indent on
set nocompatible
set autoread
set smartindent
set showmatch

" set leader key
let mapleader = ","
let maplocalleader = ","

" replace tab with 4 spaces
set tabstop=4
set softtabstop=4
set shiftwidth=4
set expandtab

" display and font
set cul
set number

" encoding
set encoding=utf-8
set termencoding=utf-8
set fileencodings=utf-8

" backup
set nobackup
set noswapfile
set nowritebackup

" to avoid short code warning
set shortmess=a
set cmdheight=2

" <F5> compile and run c/c++/bash/python/java/go
noremap <F5> :call CompileRunGcc()<CR>
func! CompileRunGcc()
    exec "w"
    if &filetype == 'c'
        exec "!gcc % -o %<"
        exec "!time ./%<"
    elseif &filetype == 'cpp'
        exec "!g++ -std=c++11 % -o %<"
        exec "!time ./%<"
    elseif &filetype == 'java'
        exec "!javac %"
        exec "!time java %<"
    elseif &filetype == 'sh'
        exec "!time bash %"
    elseif &filetype == 'python'
        exec "!time python %"
    elseif &filetype == 'go'
        exec "!time go run %"
    endif
endfunc

" arrow key is evil!!
noremap <Up> <NOP>
noremap <Down> <NOP>
noremap <Left> <NOP>
noremap <Right> <NOP>

" disable auto insert comment
autocmd FileType * setlocal formatoptions-=c formatoptions-=r formatoptions-=o
