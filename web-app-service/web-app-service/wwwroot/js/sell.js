$(document).ready(function() {
	"use strict";

	$(".drag-and-drop input").on("change", function (event) {
		console.log(event.target.value);
		event.target.parentElement.classList.add("filled");
	});
});