$(document).ready(function () {
	"use strict";
	
	var displayMax = $(".max-price-display");
	var displayMin = $(".min-price-display");
	var maxRangeSlider = $("#price-max");
	var minRangeSlider = $("#price-min");
	var maxRangeElement = maxRangeSlider.get(0);
	var minRangeElement = minRangeSlider.get(0);
	var currency = $("#currency").get(0);
	var offerSign = $("#offers-sign");

	function displayMaxPrice() {
		var sign = CurrencySigns[currency.value];
		displayMax.text(`${maxRangeElement.value} ${sign}`);
	}
	function displayMinPrice() {
		var sign = CurrencySigns[currency.value];
		displayMin.text(`${minRangeElement.value} ${sign}`);
	}

	function setShippingCurrency() {
		var sign = CurrencySigns[currency.value];
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

	var minValue = parseInt(minRangeElement.value);
	var maxValue = parseInt(maxRangeElement.value);
	var offset = (maxRangeElement.max - maxRangeElement.min) * 4 / 50;

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
	$(".custom-select").on("click", function () {
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