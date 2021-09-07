package main

const template = `<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>Verse of the Day</title>
<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@0.9.3/css/bulma.min.css">
</head>
<body>
<section class="hero is-fullheight">
<div class="hero-body">
<div class="">
<p class="title">
{{ verse }}
</p>
<p class="subtitle">
{{ verse_reference}}
</p>
</div>
</div>
</section></body>
</html>`
