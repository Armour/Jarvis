set nocompatible
syntax on
filetype plugin indent on

" set maplearder key
let mapleader = ","
let maplocalleader = ","

" basic config
set autoread
set smartindent
set showmatch

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
set fileencodings=ucs-bom,utf-8,gbk,cp936,gb2312,gb18030,big5,euc-jp,euc-kr,latin1

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

" <F6> Gdb for C,C++
noremap <F6> :call Rungdb()<CR>
func! Rungdb()
    exec "w"
    if &filetype == 'c'
        exec "!gcc % -g -o %<"
        exec "!gdb ./%<"
    elseif &filetype == 'cpp'
        exec "!g++ -std=c++11 % -g -o %<"
        exec "!gdb ./%<"
    endif
endfunc

" arrow key is evil!!
noremap <Up> <NOP>
noremap <Down> <NOP>
noremap <Left> <NOP>
noremap <Right> <NOP>

" insert one line without stay in insert mode
noremap <C-o> o<esc><CR>

" insert one space without stay in insert mode
noremap <C-i> i <esc>

" disable auto insert comment
autocmd FileType * setlocal formatoptions-=c formatoptions-=r formatoptions-=o
