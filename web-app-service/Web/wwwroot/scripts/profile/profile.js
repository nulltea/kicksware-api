function sidebarControlInit() {
	$(".profile-sidebar input[type=button]").click(function () {
		let newActive = $(`#section-${this.id}`)
		if (newActive.length) {
			$(".profile-section.active").toggleClass("active");
			newActive.toggleClass("active");
			let mode = newActive.attr("name")
			window.history.replaceState("Kicksware", `(Page ${mode})`, `/profile/${mode}`);
		}
	})
	let mode = location.pathname.split("/")[2]
	if (mode) {
		$(".profile-section.active").toggleClass("active");
		$(`section[name=${mode}]`).toggleClass("active");
	}
}

function profileFormInit(){
	let form = $(".profile-form");

	form.submit(function (event) {
		event.preventDefault();
		$.post(form.attr("action"), form.serialize(), function(response) {
			showAlert(response.result, response.message);
		});
	})
}

function showAlert(mode, message, lifetime = 5) {
	resetAlert(function () {
		$(".profile .alert-banner")
			.addClass(mode)
			.text(message)
			.addClass("active")
		clearTimeout(window.lifetimeHandler)
		window.lifetimeHandler = window.setTimeout(function () {
			resetAlert();
		}, lifetime * 1000);
	});
}


function resetAlert(callback) {
	let banner = $(".profile .alert-banner");
	if (callback) {
		if (banner.hasClass("active")){
			requestAnimationFrame(function () {
				banner.removeClass("active success error warning").text("");
			})
			window.setTimeout(callback, 500);
		}
		callback();
	} else {
		banner.removeClass("active success error warning").text("");
	}
}

$(document).ready(function () {
	"use strict";

	sidebarControlInit();

	profileFormInit();
});