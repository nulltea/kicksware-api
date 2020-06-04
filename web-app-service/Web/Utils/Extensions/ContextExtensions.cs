﻿using System;
using System.Linq;
using Core.Constants;
using Core.Entities.Users;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Http.Features;
using Web.Handlers.Authentication;
using Web.Handlers.Authorisation;

namespace Web.Utils.Extensions
{
	public static class ContextExtensions
	{
		public static AuthToken RetrieveCookieAuthToken(this HttpContext context)
		{
			// Request cookie
			if (!context.Request.Cookies.TryGetValue(MiddlewareAuthDefaults.AuthCookieName, out var protectedToken))
			{
				// If else user response set cookie header
				var setCookieHeaders = context.Response.GetTypedHeaders().SetCookie;

				if (setCookieHeaders == null) return string.Empty;

				var tokenCookie = setCookieHeaders.FirstOrDefault(cookie => cookie.Name.Equals(MiddlewareAuthDefaults.AuthCookieName));

				protectedToken = tokenCookie?.Value.Value;
			}

			if (string.IsNullOrEmpty(protectedToken)) return protectedToken;

			var unprotectedToken = protectedToken.UnprotectToken(context);

			return unprotectedToken ?? string.Empty;
		}

		internal static string GetTlsTokenBinding(this HttpContext context)
		{
			var binding = context.Features.Get<ITlsTokenBindingFeature>()?.GetProvidedTokenBindingId();// ?? Constants.AlternativeTlsBinding;
			return binding == null ? null : Convert.ToBase64String(binding);
		}

		public static bool IsLockedContext(this HttpContext context, out AuthLockedResponse response)
		{
			response = default;
			if (!context.Items.ContainsKey(AuthLockedResponse.AuthLockedKey)) return false;
			response = context.Items[AuthLockedResponse.AuthLockedKey] as AuthLockedResponse;
			return true;
		}
	}
}