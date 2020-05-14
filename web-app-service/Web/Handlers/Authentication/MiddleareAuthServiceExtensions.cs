using System;
using Core.Entities.Users;
using Microsoft.AspNetCore.Authentication;
using Microsoft.AspNetCore.Http;
using Microsoft.Extensions.DependencyInjection;
using Web.Utils.Extensions;

namespace Web.Handlers.Authentication
{
	public static class AuthExtensions
	{
		public static void StoreToken(this AuthenticationProperties properties, AuthToken token)
		{
			if (properties is null) throw new ArgumentNullException(nameof(properties));
			if (token is null) throw new ArgumentNullException(nameof(token));

			var oldTokens = properties.GetTokens();
			foreach (var t in oldTokens)
			{
				properties.Items.Remove(TokenPropertyName(token.Name));
			}

			var tokenPropertyName = TokenPropertyName(token.Name);
			properties.Items.Remove(tokenPropertyName);

			properties.Parameters[tokenPropertyName] = token;
		}

		public static AuthToken RetrieveToken(this AuthenticationProperties properties, string tokenName = default)
		{
			if (properties is null) throw new ArgumentNullException(nameof(properties));

			var tokenKey = TokenPropertyName(tokenName ?? string.Empty);
			return properties.Parameters.ContainsKey(tokenKey) ? properties.GetParameter<AuthToken>(tokenKey) : null;
		}

		public static AuthToken RetrieveToken(this AuthenticationTicket ticket, string tokenName = default)
		{
			if (ticket is null) throw new ArgumentNullException(nameof(ticket));
			if (ticket.Properties is null) throw new NullReferenceException(nameof(ticket.Properties));

			return ticket.Properties.RetrieveToken(tokenName);
		}

		public static AuthenticationProperties ToAuthProperties(this AuthToken token)
		{
			var prop = new AuthenticationProperties();
			prop.StoreToken(token);
			return prop;
		}

		public static AuthToken UnprotectToken(this string token, HttpContext context)
		{
			var secureFormat = context.RequestServices.GetService<ISecureDataFormat<AuthToken>>();
			if (secureFormat is null) throw new NullReferenceException(nameof(secureFormat));

			return secureFormat.Unprotect(token, context.GetTlsTokenBinding());
		}

		private static string TokenPropertyName(string name) => string.Join(".", MiddlewareAuthDefaults.TokenPrefix, name);
	}
}