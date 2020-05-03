function autoSliderInit() {
	let autoSlider = $(".home-auto-slider");
	if (autoSlider.length) {
		autoSlider.owlCarousel(
			{
				items: 1,
				autoplay: true,
				autoplayTimeout: 10000,
				loop: true,
				nav: false,
				smartSpeed: 1200,
				dotsSpeed: 1200,
				fluidSpeed: 1200,
			});

		/* Custom dots events */
		let dots = $('.auto-slider-dot');
		if (dots.length) {
			dots.on('click', function () {
				$('.auto-slider-dot').removeClass('active');
				$(this).addClass('active');
				autoSlider.trigger('to.owl.carousel', [$(this).index(), 1200]);
			});
		}

		/* Change active class for dots when slide changes by nav or touch */
		autoSlider.on('changed.owl.carousel', function (event) {
			$('.auto-slider-dot').removeClass('active');
			$('.auto-slider-dots-list li').eq(event.page.index).addClass('active');
		});

		// add animate.css class(es) to the elements to be animated
		function setAnimation(_elem, _InOut) {
			// Store all animationend event name in a string.
			// cf animate.css documentation
			let animationEndEvent = 'webkitAnimationEnd mozAnimationEnd MSAnimationEnd o animationend animationend';

			_elem.each(function () {
				let $elem = $(this);
				let $animationType = 'animated ' + $elem.data('animation-' + _InOut);

				$elem.addClass($animationType).one(animationEndEvent, function () {
					$elem.removeClass($animationType); // remove animate.css Class at the end of the animations
				});
			});
		}

		// Fired before current slide change
		autoSlider.on('change.owl.carousel', function (event) {
			let $currentItem = $('.auto-slider-item', autoSlider).eq(event.item.index);
			let $elemsToAnim = $currentItem.find("[data-animation-out]");
			setAnimation($elemsToAnim, 'out');
		});

		// Fired after current slide has been changed
		autoSlider.on('changed.owl.carousel', function (event) {
			let $currentItem = $('.auto-slider-item', autoSlider).eq(event.item.index);
			let $elemsToAnim = $currentItem.find("[data-animation-in]");
			setAnimation($elemsToAnim, 'in');
		})
	}
}

$(document).ready(function () {
	"use strict";

	autoSliderInit();
});