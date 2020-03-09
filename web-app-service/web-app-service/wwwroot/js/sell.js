$(document).ready(function() {
	"use strict";

	$(".drag-and-drop input").on("change", function (event) {
		console.log(event.target.parentElement);
		event.target.parentElement.classList.add("filled");
	});
});