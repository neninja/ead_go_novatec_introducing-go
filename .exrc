let g:current_filetype=""

function! s:ExecuteCurrentBufferFile()
    set errorformat=%f:%l:%c:\ %m,%f:%l:\ %m,%*\\s%f:%l\ %m,%-G%.%#

    if g:current_terminal
        " ocorre erro caso o buffer não exista mais
        try
            execute "bd! ".g:current_terminal
        catch
        endtry
    endif
    cclose

    new
    set filetype=go
    set buftype=terminal
    silent 0r!go run #
    let g:current_terminal = bufnr()
    normal! Gdd
    silent cgetbuffer
    execute "b ".g:current_terminal
    set filetype=txt
    cw
endfunction

nnoremap <leader>r :call <SID>ExecuteCurrentBufferFile()<CR>

" autocmd! BufRead *.go nnoremap <leader>r :call <SID>ExecuteCurrentBufferFile()<CR>


