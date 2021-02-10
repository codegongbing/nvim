## User Custom

Thinkvim will generate `.thinkvim.d` folder under your home path,This folder
will has two files `plugins.yaml` and `init.vim`

## Custom leaderkey and localleader key

Just config this variable

```viml
let g:mapleader
let g:maplocalleader
```

## Install plugins

I suggest that most of plugins should lazyload by dein. There has two ways to install plugins.

- if you are a vim beginner or you want to follow master branch, you should use `plugins.yaml`. like you want use choosewin plugin,you can installed it in `plugins.yaml`(notice thinkvim has installed, just for example)

```yaml
- repo: t9md/vim-choosewin
  on_map: { n: <Plug> }
  hook_source: source  $VIM_PATH/core/layers/tools/choosewin/config.vim
```

- repo means github repo
- on_map is a dein lazyload feature,that mean the plugin will loading when you trigger `on_map` key,
- hook_source is executed before plugins are sourced

Dein is hard to use for some vim beginners，you can check the `core/dein/plugins.yaml` for use.
more usage about [Dein](https://github.com/Shougo/dein.vim/blob/master/doc/dein.txt).

- if you have vimscript skill and not a vim beginner, want change this config to your own config, you can edit the `core/dein/plugins.yaml`.

the keymap of plugins you should config it in `custom.vim`, the same as installed plugins. if you follow the master branch use `custom.vim`. if you want change to your own config, edit the `core/plugins/allkey.vim`.

example:

```viml
if dein#tap('vim-choosewin')
	nmap -         <Plug>(choosewin)
	nmap <Leader>- :<C-u>ChooseWinSwapStay<CR>
endif
```

- If you want to disable some of the plugins I use, you can overwrite them, e.g.:

```
- { repo: tryu/caw.vim, if: 0 }
```

- Your plugin shortcuts should also be based on this format. used `if dein#tap('pluginname')`
- Notice ‼️ the lazyload feature of dein will not loading the default config of plugins, so you should config the keymap of lazyload plugins

register the `custom` keymap to `whichkey`, there are two settings map.

```viml
let g:which_key_map   for leaderkey
let g:which_key_localmap  for localkey
```

example:

```viml
if dein#tap('defx.nvim')
         nnoremap <silent> <Leader>F
                \ :<C-u>Defx -search=`expand('%:p')` `getcwd()`<CR>
endif
let g:which_key_map.F='test'
```

- if you edit the plugins.yaml, it doesn't work. you can try to recache the cache. use this function `call dein#recache_runtimepath()`
- you can check [my .thinkvim.d](https://github.com/taigacute/dotfiles/tree/master/thinkvim) for more usage.

## change colorscheme

- ThinkVim provide a easy way to change colorscheme, use this keymap `Leader+tc` then thinkvim will write the colorscheme that your choose into cache file,ThinkVim will read this file to load colorscheme when next startup

- Add your favorite coloscheme file into colors folder

## Spaceline

check out the [spaceline config](https://github.com/taigacute/spaceline.vim)

## Dashboard

- Custom [dashboard-nvim](https://github.com/hardcoreplayers/dashboard-nvim) header and footer

check the usage in
[dashbaord-nvim](https://github.com/hardcoreplayers/dashboard-nvim))

## Database

If you want to use the `.env` file to management your database connection.you
can use `let g:thinkvim_load_env_db =1`. then touch .env file in your homepath
Add your Database connection the prefix is `DB_CONNECTION_` like:

```
DB_CONNECTION_CRAWLER=mysql://root@localhost/crawler
```

then `Leader o d` to open database buffer. type `o` to input password .

## Custom keymap

```viml
" disable all keymaps of plugins
let g:thinkvim_disable_mappings = 0
" disable some plugins keymap is a list
let g:thinkvim_disable_pmaping = []
```

The value of `thinkvim_disable_pmaping` are

```
plugins:
dein.vim vim-buffet markdown-preview.nvim
indentLine any-jump.vim vim-floaterm
vim-dadbod-ui dash.vim coc-clap
coc.nvim vim-clap vim-go vim-delve
vimagit vim-choosewin caw.vim vim-smoothie
goyo.vim defx.nvim vim-quickrun vim-easymotion
vim-smartchr iron.nvim vim-sandwich dashboard-nvim
vim-fugitive vim-mundo vista.vim

general:
insert command quit window session toggle normal
```

eg:

```viml
let g:thinkvim_disable_pmaping = ['dein.vim','insert'....]
```

Why add this support, Because vim is free, and key mapping everyone has everyone's habits, so we can't bundle them, it should be customized by the user. more information on this issue https://github.com/hardcoreplayers/ThinkVim/issues/105.

## Lsp

- Go use gopls (need user install)
- C C++ OC use ccls (need user install)
- Python support `coc-python`,When install thinkvim if you choose python,when open python file it will auto install mpls.
- Dockerfile use docker-langserver (need user install)
- Haskell use https://github.com/haskell/haskell-ide-engine (need user install)
- Lua use https://github.com/Alloyed/lua-lsp
- Javascript Typescript use coc-tsserver
- Php use https://github.com/bmewburn/vscode-intelephense (need user install)
- Shell use https://github.com/bash-lsp/bash-language-server (need user install)
- R use coc-R
- Ruby use coc-solargraph
- Rust use coc-rust-analyzer
- Scala use coc-metals
- Vue use coc-vetur

## Add Language support

It's easy to add new language support, just run this script,then choose the new languages that you write. it will write the languages plugins to your `.thinkvim.d/plugins.yaml`

```console
cd .config/nvim/bin
./generator
```
