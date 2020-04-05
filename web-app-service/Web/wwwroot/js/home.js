/* JS Document */

/******************************

[Table of Contents]

1. Vars and Inits
3. Init Home Slider
6. Init Isotope


******************************/

$(document).ready(function () {
	"use strict";

	/* 

	1. Vars and Inits

	*/

    initHomeSlider();
    /* 

    3. Init Home Slider

    */

	function initHomeSlider() {
		if ($('.home_slider').length) {
			var homeSlider = $('.home_slider');
			homeSlider.owlCarousel(
				{
					items: 1,
					autoplay: true,
					autoplayTimeout: 10000,
					loop: true,
					nav: false,
					smartSpeed: 1200,
					dotsSpeed: 1200,
					fluidSpeed: 1200
				});

			/* Custom dots events */
			if ($('.home_slider_custom_dot').length) {
				$('.home_slider_custom_dot').on('click', function () {
					$('.home_slider_custom_dot').removeClass('active');
					$(this).addClass('active');
					homeSlider.trigger('to.owl.carousel', [$(this).index(), 1200]);
				});
			}

			/* Change active class for dots when slide changes by nav or touch */
			homeSlider.on('changed.owl.carousel', function (event) {
				$('.home_slider_custom_dot').removeClass('active');
				$('.home_slider_custom_dots li').eq(event.page.index).addClass('active');
			});

			// add animate.css class(es) to the elements to be animated
			function setAnimation(_elem, _InOut) {
				// Store all animationend event name in a string.
				// cf animate.css documentation
				var animationEndEvent = 'webkitAnimationEnd mozAnimationEnd MSAnimationEnd oanimationend animationend';

				_elem.each(function () {
					var $elem = $(this);
					var $animationType = 'animated ' + $elem.data('animation-' + _InOut);

					$elem.addClass($animationType).one(animationEndEvent, function () {
						$elem.removeClass($animationType); // remove animate.css Class at the end of the animations
					});
				});
			}

			// Fired before current slide change
			homeSlider.on('change.owl.carousel', function (event) {
				var $currentItem = $('.home_slider_item', homeSlider).eq(event.item.index);
				var $elemsToanim = $currentItem.find("[data-animation-out]");
				setAnimation($elemsToanim, 'out');
			});

			// Fired after current slide has been changed
			homeSlider.on('changed.owl.carousel', function (event) {
				var $currentItem = $('.home_slider_item', homeSlider).eq(event.item.index);
				var $elemsToanim = $currentItem.find("[data-animation-in]");
				setAnimation($elemsToanim, 'in');
			})
		}
	}
});