function initSearch() {
	if ($(".search").length && $(".search_panel").length) {
		var header = $("header");
		var search = $(".search");
		var panel = $(".search_panel");
		var dismiss = $(".search_panel .close-button");
		var searchInput = $(".search_panel input");

		search.on("click", function () {
			panel.toggleClass("active");
			setTimeout(function () {
				searchInput.focus();
			}, 500);
		});
		dismiss.on("click", function () {
			panel.toggleClass("active");
		});
		window.addEventListener("click", function (event) {
			if (panel.hasClass("active") && !isDescendant(header[0], event.target)) {
				panel.toggleClass("active");
			}
		});
	}
}

$(document).ready(function () {
	initSearch();
});