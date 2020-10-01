<div align="center">
	<img width="500" src=".github/logo.svg" alt="pinpt-logo">
</div>

<p align="center" color="#6a737d">
	<strong>This repo contains Golang utilities for working with Atlassian Confluence Wiki Format content</strong>
</p>


## Overview

The [Confluence Wiki Format](https://jira.atlassian.com/secure/WikiRendererHelpAction.jspa?section=all) represents Wiki markup used in Atlassian products. For example, in Jira Cloud platform, the text in issue descriptions from a Web Hook callback is stored in this format.

This library provides a set of utilities in Golang for dealing with Confluence Wiki formatted content.


## Install

```
go get -u github.com/pinpt/confluence
```

## Usage

import github.com/pinpt/confluence

```golang
buf, err := confluence.ParseToHTML([]byte(`||heading 1||heading 2||heading 3||
|col A1|col A2|col A3|
|col B1|col B2|col B3|`))
```

## License

Licensed under the MIT License.
