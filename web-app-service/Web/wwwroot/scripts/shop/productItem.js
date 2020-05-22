function initCarousels() {
	$(".carousel-wrapper").each(function () {
		let carousel = $(this);
		carousel.find(".flickity-button").appendTo($(this));
		carousel.find(".flickity-page-dots .dot").detach();
	})
}

function favoriteInit(){
	$(".title-header .favorite-product input[type=checkbox]").change(function () {
		let id = window.location.pathname.split("/").slice(-1)
		let checked = $(this).is(":checked");
		$.get(`/shop/${checked ? "like" : "unlike"}/${id}`);
	})
	$(".carousel-cell .favorite-product input[type=checkbox]").change(function () {
		let id = $(this).closest(".carousel-cell").attr("id")
		let checked = $(this).is(":checked");
		$.get(`/shop/${checked ? "like" : "unlike"}/${id}`);
	})
}

$(document).ready(function () {
	initCarousels();

	favoriteInit();
});