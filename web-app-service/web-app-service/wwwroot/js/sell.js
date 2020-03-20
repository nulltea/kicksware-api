$(document).ready(function() {
	"use strict";

	$(".drag-and-drop input").on("change", function (event) {
		event.target.parentElement.classList.add("filled");
	});


	var display = $(".price-display");
	var rangeSlider = $(".price-slider");
	var currency = $("#currency");

	function displayPrice() {
		var sign = CurrencySigns[currency.get(0).value];
		display.text(`${rangeSlider.get(0).value} ${sign}`);
		console.log(currency.get(0).value);
	}
	displayPrice();
	rangeSlider.on("input", displayPrice);
	$(".custom-select").on("click", displayPrice);
});