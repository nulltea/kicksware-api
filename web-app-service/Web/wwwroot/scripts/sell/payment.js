$(document).ready(function () {
	"use strict";

	let displayMax = $(".max-price-display");
	let displayMin = $(".min-price-display");
	let maxRangeSlider = $("#price-max");
	let minRangeSlider = $("#price-min");
	let maxRangeElement = maxRangeSlider.get(0);
	let minRangeElement = minRangeSlider.get(0);
	let currency = $("#currency").get(0);
	let offerSign = $("#offers-sign");

	function displayMaxPrice() {
		let sign = CurrencySigns[currency.value];
		displayMax.text(`${maxRangeElement.value} ${sign}`);
	}
	function displayMinPrice() {
		let sign = CurrencySigns[currency.value];
		displayMin.text(`${minRangeElement.value} ${sign}`);
	}

	function setShippingCurrency() {
		let sign = CurrencySigns[currency.value];
		$(".shipping-info span").each(function() {
			$(this).text(sign);
		});
	}

	function offsetEnabled() {
		if (offerSign.get(0).checked) {
			minRangeSlider.addClass("active");
			displayMin.addClass("active");
		}
	}

	function round(num) {
		return parseInt(Math.round(num * 0.2, 0) * 5);
	}

	let minValue = parseInt(minRangeElement.value);
	let maxValue = parseInt(maxRangeElement.value);
	let offset = (maxRangeElement.max - maxRangeElement.min) * 4 / 50;

	function handleCollisionMin() {
		minValue = parseInt(minRangeElement.value);
		maxValue = parseInt(maxRangeElement.value);
		offset = (maxRangeElement.max - maxRangeElement.min) * 4 / 50;
		

		if (minValue > maxValue - offset) {
			maxRangeElement.value = minValue + offset;
			displayMaxPrice();
			if (maxValue === parseInt(maxRangeElement.max)) {
				minRangeElement.value = round(parseInt(maxRangeElement.max) - offset);
			}
		}
	}
	function handleCollisionMax() {
		minValue = parseInt(minRangeElement.value);
		maxValue = parseInt(maxRangeElement.value);
		offset = (maxRangeElement.max - maxRangeElement.min) * 4 / 50;

		if (maxValue < minValue + offset) {
			minRangeElement.value = maxValue - offset;
			displayMinPrice();
			if (minValue === parseInt(minRangeElement.min)) {
				maxRangeElement.value = round(parseInt(maxRangeElement.min) + offset);
			}
		}
	}

	maxRangeSlider.on("input", function () {
		if (minRangeSlider.hasClass("active")) {
			handleCollisionMax();
		}
		displayMaxPrice();
	});
	minRangeSlider.on("input", function () {
		if (minRangeSlider.hasClass("active")) {
			handleCollisionMin();
		}
		displayMinPrice();
	});
	$(".list-select").on("click", function () {
		displayMaxPrice();
		displayMinPrice();
		setShippingCurrency();
	});
	offerSign.on("change", function () {
		minRangeSlider.toggleClass("active");
		displayMin.toggleClass("active");
	});

	displayMaxPrice();
	displayMinPrice();
	offsetEnabled();
});