var app = new Vue({
	el: '#bae-app',
	data: {
		apiAddress: this.apiAddress,
		apiPath: "songs/slug(tim-mcgraw)"
	},
	methods: {
		makeRequest: function (event) {
			var apiDemoResponse = document.getElementById("api-demo-response");
			apiDemoResponse.className = "";

			$.ajax({
				url: this.apiAddress + this.apiPath,
				method: 'GET'
			}).done(function (data) {
				document.getElementById("json-preview").textContent = JSON.stringify(data, null, 4);
				document.getElementById("api-demo-response-header").textContent = "api response - http(200)";
			}).fail(function (xhrErr, status, err) {
				document.getElementById("api-demo-response-header").textContent = "api response - http(" + xhrErr.status + ")";
				document.getElementById("json-preview").textContent = JSON.stringify(xhrErr.responseJSON, null, 4);
			}).always(function () {
				apiDemoResponse.className = " active";
				
				document.getElementById("api-demo-response").style.display = "block";
				hljs.highlightBlock($('pre code')[0]);
			});
		}
	}
});
