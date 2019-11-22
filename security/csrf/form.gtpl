<!DOCTYPE html>
<html>
<head>
<title>Page Title</title>
</head>
<body>

<form method="post" action="/validate">
    <input type="hidden" name="token" value="{{.}}"/>
    <button type="submit">Submit</button>
</form>

</body>
</html>
