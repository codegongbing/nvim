if &compatible
	" vint: -ProhibitSetNoCompatible
	set nocompatible
	" vint: +ProhibitSetNoCompatible
endif

" Set main configuration directory as parent directory
let $VIM_PATH = fnamemodify(resolve(expand('<sfile>:p')), ':h:h')
let $THINKVIM = expand($HOME.'/.thinkvim.d')

" set the user config file
let s:user_init_config = expand($THINKVIM.'/init.vim')

" 加载vim-devicons
let g:webdevicons_enable = 1
let g:webdevicons_enable_nerdtree = 1
let g:webdevicons_enable_airline_tabline = 1
let g:WebDevIconsUnicodeDecorateFileNodesDefaultSymbol = 'ƛ'

" Disable vim distribution plugins
let g:loaded_gzip = 1
let g:loaded_tar = 1
let g:loaded_tarPlugin = 1
let g:loaded_zip = 1
let g:loaded_zipPlugin = 1

let g:loaded_getscript = 1
let g:loaded_getscriptPlugin = 1
let g:loaded_vimball = 1
let g:loaded_vimballPlugin = 1

let g:loaded_matchit = 1
let g:loaded_matchparen = 1
let g:loaded_2html_plugin = 1
let g:loaded_logiPat = 1
let g:loaded_rrhelper = 1

let g:loaded_netrw = 1
let g:loaded_netrwPlugin = 1
let g:loaded_netrwSettings = 1
let g:loaded_netrwFileHandlers = 1

" Initialize base requirements
if has('vim_starting')
	" Global Mappings "{{{
	" Use spacebar as leader and ; as secondary-leader
	" Required before loading plugins!
	let g:mapleader="\<Space>"
	let g:maplocalleader=';'

	" Release keymappings prefixes, evict entirely for use of plug-ins.
	nnoremap <Space>  <Nop>
	xnoremap <Space>  <Nop>
	nnoremap ,        <Nop>
	xnoremap ,        <Nop>
	nnoremap ;        <Nop>
	xnoremap ;        <Nop>

endif

call utils#source_file($VIM_PATH,'core/packman.vim')
call utils#source_file($VIM_PATH,'core/general.vim')
call utils#source_file($VIM_PATH,'core/filetype.vim')

" Load user init config
call utils#check_source(s:user_init_config)

" disable all keymaps of plugins
let g:thinkvim_disable_mappings = 0
" disable some plugins keymap
let g:thinkvim_disable_pmaping = []

" Load key map
call utils#source_file($VIM_PATH,'keybinds/keybinds.vim')

" Initialize user favorite colorscheme
call theme#init()

if exists("*UserInit")
	call UserInit()
endif

set secure

autocmd BufWritePre *.go :silent call CocAction('runCommand', 'editor.action.organizeImport')

let g:coc_global_extensions = [
	\ 'coc-json',
	\ 'coc-vimlsp',
	\ 'coc-calc',
	\ 'coc-clangd',
	\ 'coc-css',
	\ 'coc-git',
	\ 'coc-go',
	\ 'coc-highlight',
	\ 'coc-fzf-preview',
	\ 'coc-html',
	\ 'coc-json',
	\ 'coc-markdownlint',
	\ 'coc-python',
	\ 'coc-rls',
	\ 'coc-rust-analyzer',
	\ 'coc-rome',
	\ 'coc-sh',
	\ 'coc-stylelint',
	\ 'coc-snippets',
	\ 'coc-toml',
	\ 'coc-yaml',
	\ 'coc-sql',
	\ 'coc-translator',
  \ ]
" vim: set ts=2 sw=2 tw=80 noet :

let g:buffet_powerline_separators = 1
let g:buffet_show_index = 1
let g:buffet_use_devicons = 1
let g:buffet_left_trunc_icon = "\uf0a8"
let g:buffet_right_trunc_icon = "\uf0a9"
