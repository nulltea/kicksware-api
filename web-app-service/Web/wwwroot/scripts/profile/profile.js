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
			console.log(response.success);
		});
	})
}


$(document).ready(function () {
	"use strict";

	sidebarControlInit();

	profileFormInit();
});