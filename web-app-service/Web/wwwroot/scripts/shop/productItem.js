function initCarousels() {
	$("section.carousel").each(function () {
		let carousel = $(this);
		carousel.find(".flickity-button").appendTo($(this));
		carousel.find(".flickity-page-dots .dot").detach();
	})
}

$(document).ready(function () {
	initCarousels();
});