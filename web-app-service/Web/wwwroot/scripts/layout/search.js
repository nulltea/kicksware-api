function searchPanelInit() {
	let icon = $(".search");
	let panel = $(".search-panel");
	if (icon.length && panel.length) {
		let header = $("header");
		let dismiss = $(".search-panel .close-button");
		let searchInput = $(".search-panel input");

		icon.on("click", function () {
			panel.toggleClass("active");
			setTimeout(function () {
				searchInput.focus();
			}, 1200);
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

function mainSearchInit() {
	let searchForm = $(".search-panel form[method=get]");
	let searchInput = searchForm.find("input[type=search]");
	let action = searchForm.attr("action")

	searchInput.on("input", function () {
		submitSearch(action, this.value);
	})
	$(window).keypress(function(event) {
		if (event.keyCode === 13) {
			event.preventDefault();
			if ($(event.target).is(searchInput)) {
				submitSearch(action, event.target.value);
			}
			return false;
		}
	});
}


function submitSearch(actionUrl, prefix) {
	$.get(`${actionUrl}?prefix=${prefix}`, function (response) {
		if (!response["success"]) {
			return;
		}
		$(".search-results").html(response["content"]);
	})
}


$(document).ready(function () {
	searchPanelInit();

	mainSearchInit();
});