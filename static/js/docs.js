$(document).ready(function () {
	$('.indentJson').each(function () {
		var codeElement = $(this);
		var json = codeElement.data('json');
		var jsonPretty = JSON.stringify(json, null, 2);
		codeElement.text(jsonPretty);
	});

	$('.indentHttpAndJson').each(function () {
		var codeElement = $(this);
		var json = codeElement.data('json');
		var jsonPretty = JSON.stringify(json, null, 2);

		var http = codeElement.data('http');

		if (json == "") {
			codeElement.text(http);
		} else {
			codeElement.text(http + "\n\n" + jsonPretty);
		}
	});
});
