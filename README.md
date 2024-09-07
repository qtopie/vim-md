# vim-md

*Markdown support for Vim. [plugin written in golang]*

## Features

- [x] Paste image from system clipboard and create assets automatically
- [ ] Markdown live preview
- [ ] Extended markdown syntax support when editing

## Usage

Put this in your vim plugin manager. Here's an example for `vim-plug`

```vimscript
Plug 'qtopie/vim-md', { 'do': ':VimMdUpdate' }
```

Copy image to your system clipboard, then parse it to vim with `MarkdownImagePaste`
