# Vani

`vani` is a lightweight LSP server that provides English word completions and definitions for plain text files.

---

## ✨ Features

- English word auto-completion
- Hover support to show definitions
- Works with Neovim via native LSP

---

## 🧩 Usage (Neovim)

### 1. Download or build the binary

```sh
go build -o vani main.go
# or download prebuilt binaries
```

Place it somewhere in your `$PATH` or give the absolute path in config.

---

### 2. Create LSP client config (e.g. `lua/lsp/vani.lua`)

```lua
-- lua/lsp/vani.lua

local client_id = vim.lsp.start({
  name = 'vani',
  cmd = { 'path/to/vani' }, -- replace with actual binary path
})

if not client_id then
  vim.notify("Vani LSP: failed to start client", vim.log.levels.ERROR)
  return
end

vim.api.nvim_create_autocmd("BufReadPost", {
  pattern = "text", -- or "markdown", etc.
  callback = function(args)
    local ok = vim.lsp.buf_is_attached(args.buf, client_id)
    if not ok then
      vim.lsp.buf_attach_client(args.buf, client_id)
    end
  end
})
```

---

### 3. Register it in `lspconfig`

```lua
-- plugins/lspconfig.lua
return {
  "neovim/nvim-lspconfig",
  config = function()
    -- other LSPs
    require("lsp.vani") -- path to your vani config
  end,
}
```

---

## 🗂️ TODO

- [ ] VSCode extension
- [ ] Add persistent DB for definitions

---

## 💡 Notes

- Definitions are fetched live (not persistent yet)


