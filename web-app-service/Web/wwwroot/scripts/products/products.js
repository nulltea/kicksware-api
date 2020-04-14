function initFilterPanel() {
	$(".toggle-menu").click(function () {
		let filterMenu = $(".filter-control");
		let accordion = $(".accordion");

		if (filterMenu.hasClass("active")){
			accordion.toggleClass("collapsed");
			$(".toggle-menu span").text("SHOW FILTERS");
			$(this).toggleClass("pressed");
		} else {
			$(".toggle-menu span").text("HIDE FILTERS");
			$(this).toggleClass("pressed");
			accordion.toggleClass("collapsed");
		}

		filterMenu.toggleClass("active");
	});

	$(".accordion-control").change(function (e) {
		$(".spacer").height($(".accordion").height());
	});
}

function initIsotope() {
	let sortingButtons = $('.product_sorting_btn');
	let sortNums = $('.num_sorting_btn');
	let products = $(".product_grid");
	if (products.length) {
		let grid = products.isotope({
			itemSelector: '.product',
			layoutMode: 'fitRows',
			fitRows:
				{
					gutter: 30,
				},
			getSortData:
				{
					price: function (itemElement) {
						let priceEle = $(itemElement).find('.product_price').text().replace('$', '');
						return parseFloat(priceEle);
					},
					name: '.product_name',
					stars: function (itemElement) {
						let starsElements = $(itemElement).find('.rating');
						return starsElements.attr("data-rating");
					},
				},
			animationOptions:
				{
					duration: 750,
					easing: 'linear',
					queue: false,
				},
		});

		// Sort based on the value from the sorting_type dropdown
		sortingButtons.each(function () {
			$(this).on('click',
				function () {
					let parent = $(this).parent().parent().find('.sorting_text');
					parent.text($(this).text());
					let option = $(this).attr('data-isotope-option');
					option = JSON.parse(option);
					grid.isotope(option);
				});
		});
	}
}

$(document).ready(function () {
	initFilterPanel();
});