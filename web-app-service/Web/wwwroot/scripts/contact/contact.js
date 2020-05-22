function mapInit() {
	let map;
	map = new google.maps.Map(document.getElementById('map'), {
		center: {lat: -50.415, lng: 30.635},
		zoom: 8
	});

	// Re-center map after window resize
	google.maps.event.addDomListener(window, 'resize', function()
	{
		setTimeout(function()
		{
			google.maps.event.trigger(map, "resize");
			map.setCenter(myLatlng);
		}, 1400);
	});
}

function myMap() {
	var mapProp= {
		center:new google.maps.LatLng(51.508742,-0.120850),
		zoom:5,
	};
	var map = new google.maps.Map(document.getElementById("googleMap"),mapProp);
}

$(document).ready(function () {
	"use strict";
	mapInit();
});