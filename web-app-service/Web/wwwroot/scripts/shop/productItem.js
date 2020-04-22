$(document).ready(function () {
    /* 

    5. Init Image

    */

	function initImage() {
		var images = $('.details_image_thumbnail');
		var selected = $('.details_image_large img');

		images.each(function () {
			var image = $(this);
			image.on('click', function () {
				var imagePath = new String(image.data('image'));
				selected.attr('src', imagePath);
				images.removeClass('active');
				image.addClass('active');
			});
		});
	}

	/* 

	6. Init Quantity

	*/

	function initQuantity() {
		// Handle product quantity input
		if ($('.product_quantity').length) {
			var input = $('#quantity_input');
			var incButton = $('#quantity_inc_button');
			var decButton = $('#quantity_dec_button');

			var originalVal;
			var endVal;

			incButton.on('click', function () {
				originalVal = input.val();
				endVal = parseFloat(originalVal) + 1;
				input.val(endVal);
			});

			decButton.on('click', function () {
				originalVal = input.val();
				if (originalVal > 0) {
					endVal = parseFloat(originalVal) - 1;
					input.val(endVal);
				}
			});
		}
	}

	/* 

	7. Init Isotope

	*/

	function initIsotope() {
		var sortingButtons = $('.product_sorting_btn');
		var sortNums = $('.num_sorting_btn');

		if ($('.product_grid').length) {
			var grid = $('.product_grid').isotope({
				itemSelector: '.product',
				layoutMode: 'fitRows',
				fitRows:
				{
					gutter: 30
				},
				getSortData:
				{
					price: function (itemElement) {
						var priceEle = $(itemElement).find('.product_price').text().replace('$', '');
						return parseFloat(priceEle);
					},
					name: '.product_name',
					stars: function (itemElement) {
						var starsEle = $(itemElement).find('.rating');
						var stars = starsEle.attr("data-rating");
						return stars;
					}
				},
				animationOptions:
				{
					duration: 750,
					easing: 'linear',
					queue: false
				}
			});
		}
	}
});