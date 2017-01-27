package areatz

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	// sets up a test server to serve the HTML
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(testHTML))
	}))
	areacodeURL = ts.URL

	defer ts.Close()

	assert := assert.New(t)

	codes, err := GetAreaCodes()
	json_output, err := AreaCodesToJSON()

	assert.NoError(err, "GetAreaCodes error")
	assert.NotNil(codes, "codes should not be nil")
	assert.True(len(codes) > 0)
	assert.Equal("201", codes[0].AreaCode)
	assert.Equal(-5, codes[0].GMTOffset)
	assert.Equal(true, codes[0].DST)
	assert.Equal("New Jersey", codes[0].State)
	assert.Equal("Hackensack, Jersey City, Union City, Rutherford, Leonia  ", codes[0].Region)

	assert.Equal("test", json_output[0])
}

const testHTML = `
<!DOCTYPE html>

<html xmlns="http://www.w3.org/1999/xhtml">
<head id="ctl00_Head1"><title>
	Area Code Chart with Time Zone and Current Time
</title><meta content="Get the current Time Zone (along with Current Time) for each Area Code. Includes US and Canada Area Codes." name="description" /><meta content="zip code database, zipcodes, postal, download, csv, latitude, list, commercial, locator, census data, demographics" name="keywords" /><meta http-equiv="X-UA-Compatible" content="IE=Edge" /><meta name="viewport" content="width=device-width, initial-scale=1" /><link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.2/css/bootstrap.min.css" /><link href="content/css/Site.css" rel="stylesheet" />

    <!-- added here to prevent flicker when late-loading (moved up from lower js section)-->
    <script src="//ajax.googleapis.com/ajax/libs/jquery/1.10.2/jquery.min.js" ></script>

    <script type="text/javascript">
        var pathname = window.location.pathname;
        if (pathname != "/" && pathname != "/Default.aspx") {
            document.write("<script src='//ajax.googleapis.com/ajax/libs/jqueryui/1.9.1/jquery-ui.min.js' type='text/javascript'><\/script>"); //load jquery-ui (but not on homepage)
            document.write("<link href='../../content/metro/css/jquery-ui-min.css' rel='stylesheet' type='text/css' \/>"); //todo: replace with bootstrap at some point (currently used by dialogs e.g. registerpayment and possibly elsewhere)
        }
    </script>
    <style>
        .logo {padding: 10px 70px 10px 0;}
        #loginlink {color:#fff; font-size:11pt; }
        #loginlink a {text-decoration:none; color:#fff;font-family:Roboto, sans-serif; font-size:11pt; }
        .tagline{color:#798ea0; font-size:12.5pt; margin:-12px 0 10px 0;}
        .shoutout { color:#fff; text-align:right; width:800px; float:right; margin:-40px 20px 0 0; font-style:italic;}
    </style>


    <link href='http://fonts.googleapis.com/css?family=Audiowide' rel='stylesheet' type='text/css' />
    <style>
        th {padding:5px }
        /*.blu {background-color:#90C2DB; color:#fff; font-weight:bold; text-align:center; } */
        .time {width:150px; color:lime; background-color:#174A77; font-family: 'Audiowide', cursive; text-align:center;}
        .tz {display:none;}
        .dst {display:none;}
        .options {width:100%; text-align:right; }
        .rdoCntry label{padding:0 15px 0 0;}
    </style>


</head>
<body>

    <form method="post" action="./areacodetimezone" id="aspnetForm">


<input type="hidden" name="__VIEWSTATEGENERATOR" id="__VIEWSTATEGENERATOR" value="CD9904B7" />
<input type="hidden" name="__EVENTVALIDATION" id="__EVENTVALIDATION" value="/wEdAAqfOr24ePW1P4PPadnMDP1fmD6Elsr7kzsc4yHL1RNTB7PFMGBfYw6SNx9cRJJeekDMnDt+GpKhmrU1DA1uhHDIBtIuz3pbepiQgtjGmqzPwUK1xhy4mmjxb0Ydq6MzDVbd5kPulCZaQyYzoWo7JJGj6Mz29SOvafw38I72kvKvZuMA1QSVfkDHp3/EmSPPsOhyxtn3ijATxpOpz1peAYtnvaLTusqsCsTWjKxuKANq3S9Qs3w=" />

        <div class="navbar navbar-inverse navbar-static-top">
            <div class="container">
                <div class="navbar-header">
                    <button type="button" class="navbar-toggle" data-toggle="collapse" data-target=".navbar-collapse">
                        <span class="icon-bar"></span>
                        <span class="icon-bar"></span>
                        <span class="icon-bar"></span>
                    </button>
                    <div class="hidden-xs visible-sm visible-md visible-lg">
                        <a href="./"><img src="content/images/logo.png" class="logo" alt="commercial zip code database" /></a>
                        <div class="tagline">ENTERPRISE DATABASES SINCE 1992</div>
                    </div>
                    <div class="visible-xs hidden-xm hidden-md hidden-lg"><a href="./" style="color:#fff;">GreatData.com</a></div>
                </div>

                <br />
                <div class="navbar-collapse collapse">
                    <ul class="nav navbar-nav">
			            <li><a href="zip-code-data/usa">U.S.</a></li>
			            <li><a href="zip-code-data/canada">Canadian</a></li>
			            <li><a href="zip-code-data/mexico">Mexican</a></li>

                        <li class="dropdown">
                          <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false">More <span class="caret"></span></a>
                          <ul class="dropdown-menu" role="menu">
                            <li><a id="ctl00_hypContact" href="ContactUs.aspx">Contact</a></li>
                            <li><a id="ctl00_hypAboutUS" href="History.aspx">About</a></li>
                            <li><a id="ctl00_hypFAQ" href="Faq.aspx">FAQ</a></li>
                            <li><a id="ctl00_hpyFree" href="promo-main">Free Products</a></li>
                          </ul>
                        </li>

                        <li class="dropdown">
                          <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false">Subscribers <span class="caret"></span></a>
                          <ul class="dropdown-menu" role="menu">
                            <li><a id="ctl00_hypDownload" href="Secure/ActiveDownloadList.aspx">Download Data</a></li>
					        <li><a id="ctl00_hypAcctMaint" href="Secure/RegisterPayment.aspx">Payment Info</a></li>
                            <li><a id="ctl00_hypAddContacts" href="Secure/ContactList.aspx">Manage Contacts</a></li>
					        <li><a id="ctl00_hypActSubsc" href="Secure/ActiveSubscription.aspx">View Subscriptions</a></li>
                            <li><a id="ctl00_hypLogin" href="Login.aspx?type=manage">Login Admin</a></li>
                            <li><a id="ctl00_hypOrdHist" href="Secure/OrderHistory.aspx">Order History</a></li>
                          </ul>
                        </li>

                        <li><a href="Cart.aspx" title="cart / checkout"><img src="content/images/cart_wh.png" alt="cart data" width="22" height="14" /></a></li>
                    </ul>


                    <ul class="nav navbar-nav pull-right">
                      <li class="dropdown pull-right">
                        <a class="dropdown-toggle" href="#" data-toggle="dropdown" role="button" aria-expanded="false">
                            <div id="mnuLoginLink" class="text-right">
                                <span id="ctl00_loginInvite">Sign In</span> <!-- hidden after logged in -->
		                         <span class="caret"></span>
	                        </div>
                        </a>

                        <div class="dropdown-menu" style="padding:15px;">
                            <a id="ctl00_loginSatus" href="javascript:__doPostBack(&#39;ctl00$loginSatus$ctl02&#39;,&#39;&#39;)">Sign In</a>
                            <br />

                                    <div class="form-group">
                                        <label for="UserName">Username:</label> &nbsp;
                                        <input name="ctl00$mnuLogin$UserName" type="text" id="ctl00_mnuLogin_UserName" class="form-control" />
                                    </div>
                                    <div class="form-group">
                                        <label for="Password">Password:</label> &nbsp;
                                        <input name="ctl00$mnuLogin$Password" type="password" id="ctl00_mnuLogin_Password" class="form-control" />
                                    </div>

                                    <input type="submit" name="ctl00$mnuLogin$LoginButton" value="Log In" onclick="javascript:WebForm_DoPostBackWithOptions(new WebForm_PostBackOptions(&quot;ctl00$mnuLogin$LoginButton&quot;, &quot;&quot;, true, &quot;LoginUserValidationGroup&quot;, &quot;&quot;, false, false))" id="ctl00_mnuLogin_LoginButton" class="btn btn-primary btn-lg btn-block" />
                                    <br />




                        </div>
                      </li>
                    </ul>


                </div>

                <div class="shoutout">
                    <div id="quoteslide" class="carousel slide hidden-xs hidden-sm hidden-md">
                        <div class="carousel-inner"></div>
                    </div> <!-- / quoteslide -->
                </div>



            </div>
        </div>



	    <div id="mastercontent" class="container body-content">

    <h1>Area Code Time Zone Chart</h1>
    <div class="options">
        <a class="hidden-xs" href="../pdf/ac-tz.pdf">Printable Area Code Time Zone Chart (pdf)</a><br />

        <span id="ctl00_cph_body_rdoCountry" class="rdoCntry"><input id="ctl00_cph_body_rdoCountry_0" type="radio" name="ctl00$cph_body$rdoCountry" value="USA" checked="checked" /><label for="ctl00_cph_body_rdoCountry_0">USA</label><input id="ctl00_cph_body_rdoCountry_1" type="radio" name="ctl00$cph_body$rdoCountry" value="CAN" onclick="javascript:setTimeout(&#39;__doPostBack(\&#39;ctl00$cph_body$rdoCountry$1\&#39;,\&#39;\&#39;)&#39;, 0)" /><label for="ctl00_cph_body_rdoCountry_1">Canada</label></span>
    </div>
    <input name="ctl00$cph_body$hidClientTimeOffset" type="hidden" id="ctl00_cph_body_hidClientTimeOffset" />
    <input name="ctl00$cph_body$hidClientTime" type="hidden" id="ctl00_cph_body_hidClientTime" />



    <div>
	<table class="gvGrid" cellspacing="0" rules="all" border="1" id="ctl00_cph_body_gvTZAC">
		<tr>
			<th scope="col">Area Code</th><th class="tz" scope="col">Offset</th><th class="dst" scope="col">DSTflag</th><th scope="col">Current Time</th><th class="hidden-xs" scope="col">Std Time Zone</th><th scope="col">State/Prov.</th><th class="hidden-xs" scope="col">Region</th>
		</tr><tr>
			<td>201</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>New Jersey</td><td class="hidden-xs">Hackensack, Jersey City, Union City, Rutherford, Leonia  </td>
		</tr><tr>
			<td>202</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>District of Columbia</td><td class="hidden-xs">Washington, District of Columbia  </td>
		</tr><tr>
			<td>203</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Connecticut</td><td class="hidden-xs">Bridgeport, New Haven, Stamford, Waterbury, Norwalk, Danbury, Greenwich  </td>
		</tr><tr>
			<td>205</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Alabama</td><td class="hidden-xs">Birmingham, Huntsville, Tuscaloosa, Anniston  </td>
		</tr><tr>
			<td>206</td><td class="tz">-8</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">PST</td><td>Washington</td><td class="hidden-xs">Seattle, Everett  </td>
		</tr><tr>
			<td>207</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Maine</td><td class="hidden-xs">Maine: All regions  </td>
		</tr><tr>
			<td>208*</td><td class="tz">-7</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">MST</td><td>Idaho</td><td class="hidden-xs">Idaho: All regions *(Area Code crosses into PST time zone) </td>
		</tr><tr>
			<td>209</td><td class="tz">-8</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">PST</td><td>California</td><td class="hidden-xs">Stockton, Modesto, Merced, Oakdale  </td>
		</tr><tr>
			<td>210</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Texas</td><td class="hidden-xs">San Antonio  </td>
		</tr><tr>
			<td>212</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>New York</td><td class="hidden-xs">New York City (Manhattan only)  </td>
		</tr><tr>
			<td>213</td><td class="tz">-8</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">PST</td><td>California</td><td class="hidden-xs">Los Angeles, Compton  </td>
		</tr><tr>
			<td>214</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Texas</td><td class="hidden-xs">Dallas  </td>
		</tr><tr>
			<td>215</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Pennsylvania</td><td class="hidden-xs">Philadelphia, Lansdale, Doylestown, Newtown, Quakertown  </td>
		</tr><tr>
			<td>216</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Ohio</td><td class="hidden-xs">Cleveland, Terrace, Independence, Montrose  </td>
		</tr><tr>
			<td>217</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Illinois</td><td class="hidden-xs">Springfield, Champaign Urbana, Decatur, Central Illinois  </td>
		</tr><tr>
			<td>218</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Minnesota</td><td class="hidden-xs">Duluth, Virginia, Moorhead, Brainerd, Wadena  </td>
		</tr><tr>
			<td>219*</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Indiana</td><td class="hidden-xs">Gary, Hammond, Merrillville, Portage, Michigan City, Valparaiso *(Area Code crosses into EST time zone) </td>
		</tr><tr>
			<td>220</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Ohio</td><td class="hidden-xs">Overlay of 740 Area Code.  </td>
		</tr><tr>
			<td>224</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Illinois</td><td class="hidden-xs">Overlay of 847 area code (Chicago)  </td>
		</tr><tr>
			<td>225</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Louisiana</td><td class="hidden-xs">Baton Rouge and Surrounding Areas  </td>
		</tr><tr>
			<td>228</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Mississippi</td><td class="hidden-xs">Gulfport, Biloxi, Pascagoula, Bay St. Louis  </td>
		</tr><tr>
			<td>229</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Georgia</td><td class="hidden-xs">Albany, Valdosta, Thomasville, Bainbridge, Tifton, Americus, Moultrie, Cordele  </td>
		</tr><tr>
			<td>231</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Michigan</td><td class="hidden-xs">Muskegon, Traverse City, Big Rapids, Cadillac, Cheboygan  </td>
		</tr><tr>
			<td>234</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Ohio</td><td class="hidden-xs">Overlay of 330 Area Code  </td>
		</tr><tr>
			<td>239</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Florida</td><td class="hidden-xs">Ft. Myers, Naples, Cape Coral, Bonita Springs, Immokalee, Lehigh Acres, Sanibel, Captiva, Pine Island  </td>
		</tr><tr>
			<td>240</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Maryland</td><td class="hidden-xs">Overlay of 301 Area Code (Maryland)  </td>
		</tr><tr>
			<td>248</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Michigan</td><td class="hidden-xs">Troy, Pontiac, Royal Oak, Birmingham, Rochester, Farmington Hills  </td>
		</tr><tr>
			<td>251</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Alabama</td><td class="hidden-xs">Mobile, Prichard, Tillmans Corner, Fairhope, Jackson, Gulfshores  </td>
		</tr><tr>
			<td>252</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>North Carolina</td><td class="hidden-xs">Greenville, Rocky Mount, Wilson, New Bern  </td>
		</tr><tr>
			<td>253</td><td class="tz">-8</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">PST</td><td>Washington</td><td class="hidden-xs">Tacoma, Kent, Auburn  </td>
		</tr><tr>
			<td>254</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Texas</td><td class="hidden-xs">Waco, Killeen, Temple  </td>
		</tr><tr>
			<td>256</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Alabama</td><td class="hidden-xs">Huntsville, Anniston, Decatur, Gadsden, Florence  </td>
		</tr><tr>
			<td>260</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Indiana</td><td class="hidden-xs">Fort Wayne, Huntington, Wabash, Lagrange  </td>
		</tr><tr>
			<td>262</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Wisconsin</td><td class="hidden-xs">Green Bay, Appleton, Racine, Kenosha, Oshkosh, Waukesha, Menomonee Falls, West Bend, Sheboygan  </td>
		</tr><tr>
			<td>267</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Pennsylvania</td><td class="hidden-xs">Overlay of 215 Area Code (Philadelphia Area)  </td>
		</tr><tr>
			<td>269</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Michigan</td><td class="hidden-xs">Kalamazoo, Battle Creek, St. Joseph, Three Rivers, South Haven, Benton Harbor, Sturgis, Hastings  </td>
		</tr><tr>
			<td>270*</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Kentucky</td><td class="hidden-xs">Bowling Green, Paducah, Owensboro, Hopkinsville *(Area Code crosses into EST time zone) </td>
		</tr><tr>
			<td>272</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Pennsylvania</td><td class="hidden-xs">Overlay of 570 Area Code.  </td>
		</tr><tr>
			<td>276</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Virginia</td><td class="hidden-xs">Martinsville, Abingdon, Wytheville, Bristol, Marion, Collinsville  </td>
		</tr><tr>
			<td>281</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Texas</td><td class="hidden-xs">Houston, Sugar Land, Buffalo, Airline, Greenspoint, Spring  </td>
		</tr><tr>
			<td>301</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Maryland</td><td class="hidden-xs">Rockville, Silver Spring, Bethesda, Gaithersburg, Frederick, Laurel, Hagerstown  </td>
		</tr><tr>
			<td>302</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Delaware</td><td class="hidden-xs">Delaware: All regions  </td>
		</tr><tr>
			<td>303</td><td class="tz">-7</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">MST</td><td>Colorado</td><td class="hidden-xs">Denver, Littleton, Englewood, Arvada, Boulder, Aurora  </td>
		</tr><tr>
			<td>304</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>West Virginia</td><td class="hidden-xs">West Virginia: All regions  </td>
		</tr><tr>
			<td>305</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Florida</td><td class="hidden-xs">Miami, Perrine, Homestead, Florida Keys  </td>
		</tr><tr>
			<td>307</td><td class="tz">-7</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">MST</td><td>Wyoming</td><td class="hidden-xs">Wyoming: All regions  </td>
		</tr><tr>
			<td>308*</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Nebraska</td><td class="hidden-xs">Grand Island, Kearney, North Platte, Scottsbluff *(Area Code crosses into MST time zone) </td>
		</tr><tr>
			<td>309</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Illinois</td><td class="hidden-xs">Peoria, Bloomington, Rock Island, Galesburg, Macomb  </td>
		</tr><tr>
			<td>310</td><td class="tz">-8</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">PST</td><td>California</td><td class="hidden-xs">Compton, Santa Monica, Beverly Hills, West LA, Inglewood, Redondo, El Segundo, Culver City, Torrance  </td>
		</tr><tr>
			<td>312</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Illinois</td><td class="hidden-xs">Chicago (Downtown area), Wheeling  </td>
		</tr><tr>
			<td>313</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Michigan</td><td class="hidden-xs">Detroit, Livonia, Dearborn  </td>
		</tr><tr>
			<td>314</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Missouri</td><td class="hidden-xs">Saint Louis, Ladue, Kirkwood, Creve Coeur, Overland, Ferguson  </td>
		</tr><tr>
			<td>315</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>New York</td><td class="hidden-xs">Syracuse, Utica, Watertown, Rome  </td>
		</tr><tr>
			<td>316</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Kansas</td><td class="hidden-xs">Wichita, Hutchinson  </td>
		</tr><tr>
			<td>317</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Indiana</td><td class="hidden-xs">Indianapolis, Carmel, Fishers, Greenwood  </td>
		</tr><tr>
			<td>318</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Louisiana</td><td class="hidden-xs">Shreveport, Monroe, Alexandria, Ruston, Natchitoches   </td>
		</tr><tr>
			<td>319</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Iowa</td><td class="hidden-xs">Cedar Rapids, Iowa City, Waterloo, Burlington  </td>
		</tr><tr>
			<td>320</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Minnesota</td><td class="hidden-xs">St. Cloud, Alexandria, Willmar, Little Falls  </td>
		</tr><tr>
			<td>321</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Florida</td><td class="hidden-xs">Brevard County (Cocoa, Melbourne, Eau Gallie, Titusville), Overlay of most of 407 Area Code.  </td>
		</tr><tr>
			<td>323</td><td class="tz">-8</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">PST</td><td>California</td><td class="hidden-xs">Los Angeles, Montebello  </td>
		</tr><tr>
			<td>325</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Texas</td><td class="hidden-xs">Albilene, San Angelo, Brownwood, Synder, Swee  </td>
		</tr><tr>
			<td>330</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Ohio</td><td class="hidden-xs">Akron, Youngstown, Canton, Warren, Kent, Alliance, Medina, New Philadelphia  </td>
		</tr><tr>
			<td>331</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Illinois</td><td class="hidden-xs">Overlay of the 630 area code.  </td>
		</tr><tr>
			<td>334*</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Alabama</td><td class="hidden-xs">Montgomery, Dothan, Auburn, Selma, Opelika, Phenix City, Tuskegee *(Area Code crosses into EST time zone) </td>
		</tr><tr>
			<td>336</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>North Carolina</td><td class="hidden-xs">Greensboro, Winston Salem, Highpoint, Burlington, Lexington, Asheboro, Reidsville  </td>
		</tr><tr>
			<td>337</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Louisiana</td><td class="hidden-xs">Lake Charles, Lafayette, New Iberia, Leesville, Opelousas, Crowley  </td>
		</tr><tr>
			<td>339</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Massachusetts</td><td class="hidden-xs">Overlay of 781 area code.  </td>
		</tr><tr>
			<td>346</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Texas</td><td class="hidden-xs">Overlay of 713, 281 and 832 area codes (Houston).  </td>
		</tr><tr>
			<td>347</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>New York</td><td class="hidden-xs">Overlay of 718 Area Code  (Bronx, Brooklyn, Queens, Staten Island).  </td>
		</tr><tr>
			<td>351</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Massachusetts</td><td class="hidden-xs">Overlay of 978 area code.  </td>
		</tr><tr>
			<td>352</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Florida</td><td class="hidden-xs">Gainesville, Ocala, Leesburg, Brookville  </td>
		</tr><tr>
			<td>360</td><td class="tz">-8</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">PST</td><td>Washington</td><td class="hidden-xs">Vancouver, Olympia, Bellingham, Silverdale  </td>
		</tr><tr>
			<td>361</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Texas</td><td class="hidden-xs">Corpus Christi, Victoria  </td>
		</tr><tr>
			<td>364*</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Kentucky</td><td class="hidden-xs">Overlay of 270 Area Code. *(Area Code crosses into EST time zone) </td>
		</tr><tr>
			<td>380</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Ohio</td><td class="hidden-xs">Overlay of 614 Area Code  </td>
		</tr><tr>
			<td>385</td><td class="tz">-7</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">MST</td><td>Utah</td><td class="hidden-xs">Overlay of 801 Area Code.  </td>
		</tr><tr>
			<td>386</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Florida</td><td class="hidden-xs">Daytona, Deland, Lake City, DeBary, Orange City, New Smyrna Beach, Palatka, Palm Coast  </td>
		</tr><tr>
			<td>401</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Rhode Island</td><td class="hidden-xs">Rhode Island: All regions  </td>
		</tr><tr>
			<td>402*</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Nebraska</td><td class="hidden-xs">Lincoln, Omaha, Norfolk, Columbus *(Area Code crosses into MST time zone) </td>
		</tr><tr>
			<td>404</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Georgia</td><td class="hidden-xs">Atlanta  </td>
		</tr><tr>
			<td>405</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Oklahoma</td><td class="hidden-xs">Oklahoma City, Norman, Stillwater, Britton, Bethany, Moore  </td>
		</tr><tr>
			<td>406</td><td class="tz">-7</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">MST</td><td>Montana</td><td class="hidden-xs">Montana: All regions  </td>
		</tr><tr>
			<td>407</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Florida</td><td class="hidden-xs">Orlando, Winter Park, Kissimmee, Cocoa, Lake Buena Vista, Melbourne  </td>
		</tr><tr>
			<td>408</td><td class="tz">-8</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">PST</td><td>California</td><td class="hidden-xs">San Jose, Sunnyvale, Campbell Los Gatos, Salinas, San Martin, Saratoga, Morgan Hill, Gilroy  </td>
		</tr><tr>
			<td>409</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Texas</td><td class="hidden-xs">Galveston, Beaumont, Port Arthur, Texas City, Nederland  </td>
		</tr><tr>
			<td>410</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Maryland</td><td class="hidden-xs">Baltimore, Annapolis, Towson, Catonsville, Glen Burnie  </td>
		</tr><tr>
			<td>412</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Pennsylvania</td><td class="hidden-xs">Pittsburgh  </td>
		</tr><tr>
			<td>413</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Massachusetts</td><td class="hidden-xs">Springfield, Pittsfield, Holyoke, Amherst  </td>
		</tr><tr>
			<td>414</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Wisconsin</td><td class="hidden-xs">Milwaukee, Greendale, Franklin, Cudahy, St. Francis, Brown Deer, Whitefish Bay  </td>
		</tr><tr>
			<td>415</td><td class="tz">-8</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">PST</td><td>California</td><td class="hidden-xs">San Francisco  </td>
		</tr><tr>
			<td>417</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Missouri</td><td class="hidden-xs">Joplin, Springfield, Branson, Lebanon  </td>
		</tr><tr>
			<td>419</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Ohio</td><td class="hidden-xs">Toledo, Lima, Mansfield, Sandusky, Findlay  </td>
		</tr><tr>
			<td>423*</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Tennessee</td><td class="hidden-xs">Chattanooga, Cleveland, Johnson City, Bristol, Kingsport, Athens, Morristown, Greeneville *(Area Code crosses into CST time zone) </td>
		</tr><tr>
			<td>424</td><td class="tz">-8</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">PST</td><td>California</td><td class="hidden-xs">Overlay of 310 Area Code  </td>
		</tr><tr>
			<td>425</td><td class="tz">-8</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">PST</td><td>Washington</td><td class="hidden-xs">Bothell, Everett, Bellevue, Kirkland, Renton  </td>
		</tr><tr>
			<td>430</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Texas</td><td class="hidden-xs">Overlay of 903 Area Code.  </td>
		</tr><tr>
			<td>432</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Texas</td><td class="hidden-xs">Midland, Terminal, Odessa, Big Spring, Alpine, Pecos, Fort Stockton  </td>
		</tr><tr>
			<td>434</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Virginia</td><td class="hidden-xs">Lynchburg, Danville, Charlottesville, Madison Heights, South Boston  </td>
		</tr><tr>
			<td>435</td><td class="tz">-7</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">MST</td><td>Utah</td><td class="hidden-xs">Logan, St. George, Park City, Tooele, Brigham City, Richfield, Moab, Blanding  </td>
		</tr><tr>
			<td>440</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Ohio</td><td class="hidden-xs">Willoughby, Hillcrest, Trinity, Lorain, Elyria  </td>
		</tr><tr>
			<td>442</td><td class="tz">-8</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">PST</td><td>California</td><td class="hidden-xs">Overlay of 760 area code.  </td>
		</tr><tr>
			<td>443</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Maryland</td><td class="hidden-xs">Overlay of 410 Area Code  </td>
		</tr><tr>
			<td>456</td><td class="tz">0</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center"></td><td></td><td class="hidden-xs">Used for premium-rate calling  </td>
		</tr><tr>
			<td>458*</td><td class="tz">-8</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">PST</td><td>Oregon</td><td class="hidden-xs">Overlay of 541 Area Code. *(Area Code crosses into MST time zone) </td>
		</tr><tr>
			<td>469</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Texas</td><td class="hidden-xs">Overlay of 214, 972 Area Codes (Greater Dallas Area).  </td>
		</tr><tr>
			<td>470</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Georgia</td><td class="hidden-xs">Overlay of 404, 678, 770.  </td>
		</tr><tr>
			<td>475</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Connecticut</td><td class="hidden-xs">Overlay of 203 area code.  </td>
		</tr><tr>
			<td>478</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Georgia</td><td class="hidden-xs">Macon, Warner Robbins, Dublin, Milledgeville, Forsyth  </td>
		</tr><tr>
			<td>479</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Arkansas</td><td class="hidden-xs">Fort Smith, Fayetteville, Rogers, Bentonville, Russellville  </td>
		</tr><tr>
			<td>480</td><td class="tz">-7</td><td class="dst">N</td><td class="time"></td><td class="hidden-xs" align="center">MST</td><td>Arizona</td><td class="hidden-xs">Mesa, Scottsdale, Tempe, Chandler, Gilbert  </td>
		</tr><tr>
			<td>484</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Pennsylvania</td><td class="hidden-xs">Overlay of 610 Area Code (Allentown, Reading, Bethlehem Area)  </td>
		</tr><tr>
			<td>500</td><td class="tz">0</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center"></td><td></td><td class="hidden-xs">Used for Personal Communications Services  </td>
		</tr><tr>
			<td>501</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Arkansas</td><td class="hidden-xs">Little Rock, Hot Springs  </td>
		</tr><tr>
			<td>502</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Kentucky</td><td class="hidden-xs">Louisville, Frankfort, Fort Knox, Pleasure Ridge Park  </td>
		</tr><tr>
			<td>503</td><td class="tz">-8</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">PST</td><td>Oregon</td><td class="hidden-xs">Portland, Salem, Beaverton, Gresham, Hillsboro  </td>
		</tr><tr>
			<td>504</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Louisiana</td><td class="hidden-xs">New Orleans, Metairie, Kenner, Chalmette  </td>
		</tr><tr>
			<td>505</td><td class="tz">-7</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">MST</td><td>New Mexico</td><td class="hidden-xs">New Mexico: Albuquerque, Bernalillo, Farmington, Gallup, Grants, Las Vegas, Los Alamos, Rio Rancho, Santa Fe  </td>
		</tr><tr>
			<td>507</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Minnesota</td><td class="hidden-xs">Rochester, Mankato, Winona, Faribault, Luverne  </td>
		</tr><tr>
			<td>508</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Massachusetts</td><td class="hidden-xs">Worcester, Framingham, Brockton, Plymouth, New Bedford, Marlboro, Natick, Taunton, Auburn, Westboro, Easton  </td>
		</tr><tr>
			<td>509</td><td class="tz">-8</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">PST</td><td>Washington</td><td class="hidden-xs">Spokane, Yakima, Walla Walla, Pullman, Kenwick  </td>
		</tr><tr>
			<td>510</td><td class="tz">-8</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">PST</td><td>California</td><td class="hidden-xs">Oakland, Fremont, Newark, Hayward, Richmond  </td>
		</tr><tr>
			<td>512</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Texas</td><td class="hidden-xs">Austin, San Marcos, Round Rock, Dripping Springs   </td>
		</tr><tr>
			<td>513</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Ohio</td><td class="hidden-xs">Cincinnati, Hamilton, Clermont, Middleton, Mason  </td>
		</tr><tr>
			<td>515</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Iowa</td><td class="hidden-xs">Des Moines, Ames, Jefferson  </td>
		</tr><tr>
			<td>516</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>New York</td><td class="hidden-xs">Nassau County: Hempstead, Oceanside, Freeport, Long Beach, Garden City, Glen Cove, Mineola  </td>
		</tr><tr>
			<td>517</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Michigan</td><td class="hidden-xs">Lansing, Bay City, Jackson, Howell, Adrian  </td>
		</tr><tr>
			<td>518</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>New York</td><td class="hidden-xs">Albany, Schenectady, Troy, Glens Falls, Saratoga Springs, Lake Placid  </td>
		</tr><tr>
			<td>520</td><td class="tz">-7</td><td class="dst">N</td><td class="time"></td><td class="hidden-xs" align="center">MST</td><td>Arizona</td><td class="hidden-xs">Tucson, Sierra Vista, Nogales, Douglass, Bisbee  </td>
		</tr><tr>
			<td>530</td><td class="tz">-8</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">PST</td><td>California</td><td class="hidden-xs">Chico, Redding, Marysville, Auburn, Davis, Placerville  </td>
		</tr><tr>
			<td>531*</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Nebraska</td><td class="hidden-xs">Overlay of 402 Area Code. *(Area Code crosses into MST time zone) </td>
		</tr><tr>
			<td>533</td><td class="tz">0</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center"></td><td></td><td class="hidden-xs">Used for Personal Communications Services  </td>
		</tr><tr>
			<td>534</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Wisconsin</td><td class="hidden-xs">Overlay of 715 area code.  </td>
		</tr><tr>
			<td>539</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Oklahoma</td><td class="hidden-xs">Overlay of 918 area code.  </td>
		</tr><tr>
			<td>540</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Virginia</td><td class="hidden-xs">Roanoke, Harrisonburg, Fredericksburg, Blacksburg, Winchester, Staunton, Culpeper  </td>
		</tr><tr>
			<td>541*</td><td class="tz">-8</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">PST</td><td>Oregon</td><td class="hidden-xs">Eugene, Medford, Corvallis, Bend, Albany, Roseburg, Klamath Falls *(Area Code crosses into MST time zone) </td>
		</tr><tr>
			<td>544</td><td class="tz">0</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center"></td><td></td><td class="hidden-xs">Used for Personal Communications Services  </td>
		</tr><tr>
			<td>551</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>New Jersey</td><td class="hidden-xs">Overlay of 201 area code.  </td>
		</tr><tr>
			<td>559</td><td class="tz">-8</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">PST</td><td>California</td><td class="hidden-xs">Fresno, Visalia, Madera, San Joaquin, Porterville, Hanford  </td>
		</tr><tr>
			<td>561</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Florida</td><td class="hidden-xs">Palm Beaches, Boca Raton, Boynton Beach  </td>
		</tr><tr>
			<td>562</td><td class="tz">-8</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">PST</td><td>California</td><td class="hidden-xs">Long Beach, Norwalk, Alamitos, Downey, Whittier, Lakewood, La Habra  </td>
		</tr><tr>
			<td>563</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Iowa</td><td class="hidden-xs">Davenport, Dubuque, Clinton, Muscatine, Decorah, Manchester, West Union  </td>
		</tr><tr>
			<td>566</td><td class="tz">0</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center"></td><td></td><td class="hidden-xs">Used for Personal Communications Services  </td>
		</tr><tr>
			<td>567</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Ohio</td><td class="hidden-xs">Overlay of 419 area code.  </td>
		</tr><tr>
			<td>570</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Pennsylvania</td><td class="hidden-xs">Wilkes-Barre, Scranton, Stroudsburg, Williamsport, Pittston  </td>
		</tr><tr>
			<td>571</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Virginia</td><td class="hidden-xs">Overlay of 703 area code.  </td>
		</tr><tr>
			<td>573</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Missouri</td><td class="hidden-xs">Columbia, Cape Girardeau, Jefferson City  </td>
		</tr><tr>
			<td>574*</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Indiana</td><td class="hidden-xs">South Bend, Elkhart, Mishawaka, Granger, La Porte *(Area Code crosses into CST time zone) </td>
		</tr><tr>
			<td>575</td><td class="tz">-7</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">MST</td><td>New Mexico</td><td class="hidden-xs">Split of 505 area code. Alamogordo, Carlsbad, Clovis, Deming, El Rito, Galina, Hatch, Hobbs, Las Cruces, Penasco, Raton, Taos  </td>
		</tr><tr>
			<td>580</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Oklahoma</td><td class="hidden-xs">Lawton, Enid, Ponca City, Ardmore, Duncan  </td>
		</tr><tr>
			<td>585</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>New York</td><td class="hidden-xs">Rochester, East Rochester, Olean, Batavia, Webster, Fairport, Henrietta  </td>
		</tr><tr>
			<td>586</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Michigan</td><td class="hidden-xs">Warren, Mount Clemens, Roseville, Center Line, Utica, Romeo, Richmond, Washington, New Baltimore  </td>
		</tr><tr>
			<td>588</td><td class="tz">0</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center"></td><td></td><td class="hidden-xs">Used for Personal Communications Services  </td>
		</tr><tr>
			<td>601</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Mississippi</td><td class="hidden-xs">Southern Mississippi: Jackson, Hattiesburg, Vicksburg, Meridian  </td>
		</tr><tr>
			<td>602</td><td class="tz">-7</td><td class="dst">N</td><td class="time"></td><td class="hidden-xs" align="center">MST</td><td>Arizona</td><td class="hidden-xs">Phoenix  </td>
		</tr><tr>
			<td>603</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>New Hampshire</td><td class="hidden-xs">New Hampshire: All regions  </td>
		</tr><tr>
			<td>605*</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>South Dakota</td><td class="hidden-xs">South Dakota: All regions *(Area Code crosses into MST time zone) </td>
		</tr><tr>
			<td>606*</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Kentucky</td><td class="hidden-xs">Ashland, Winchester, Pikeville, Somerset *(Area Code crosses into CST time zone) </td>
		</tr><tr>
			<td>607</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>New York</td><td class="hidden-xs">Elmira, Ithaca, Stamford, Binghamton, Endicott, Oneonta  </td>
		</tr><tr>
			<td>608</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Wisconsin</td><td class="hidden-xs">Madison, La Crosse, Janesville, Middleton, Monroe, Platteville, Dodgeville  </td>
		</tr><tr>
			<td>609</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>New Jersey</td><td class="hidden-xs">Atlantic City, Trenton, Princeton, Pleasantville, Fort Dix, Lawrenceville  </td>
		</tr><tr>
			<td>610</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Pennsylvania</td><td class="hidden-xs">Allentown, Reading, Bethlehem, West Chester, Pottstown  </td>
		</tr><tr>
			<td>611</td><td class="tz">0</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center"></td><td></td><td class="hidden-xs">Used for telephone repair service in many Areas  </td>
		</tr><tr>
			<td>612</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Minnesota</td><td class="hidden-xs">Minneapolis, Richfield  </td>
		</tr><tr>
			<td>614</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Ohio</td><td class="hidden-xs">Columbus, Worthington, Dublin, Reynoldsburg, Westerville, Gahanna, Hilliard  </td>
		</tr><tr>
			<td>615</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Tennessee</td><td class="hidden-xs">Nashville, Mufreesboro, Hendersonville, Frank  </td>
		</tr><tr>
			<td>616</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Michigan</td><td class="hidden-xs">Grand Rapids, Holland, Grand Haven, Greenville, Zeeland, Ionia  </td>
		</tr><tr>
			<td>617</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Massachusetts</td><td class="hidden-xs">Boston, Cambridge, Quincy, Newton, Brookline, Brighton, Somerville, Dorchester, Hyde Park, Jamaica Plain  </td>
		</tr><tr>
			<td>618</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Illinois</td><td class="hidden-xs">Collinsville, Alton, Carbondale, Belleville, Mount Vernon, Centralia  </td>
		</tr><tr>
			<td>619</td><td class="tz">-8</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">PST</td><td>California</td><td class="hidden-xs">San Diego, Chula Vista, El Cajon, La Mesa, National City, Coronado  </td>
		</tr><tr>
			<td>620*</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Kansas</td><td class="hidden-xs">Hutchinson, Emporia, Liberal, Pittsburg, Dodge City, Garden City, Coffeyville *(Area Code crosses into MST time zone) </td>
		</tr><tr>
			<td>623</td><td class="tz">-7</td><td class="dst">N</td><td class="time"></td><td class="hidden-xs" align="center">MST</td><td>Arizona</td><td class="hidden-xs">Glendale, Sun City, Peoria  </td>
		</tr><tr>
			<td>626</td><td class="tz">-8</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">PST</td><td>California</td><td class="hidden-xs">Pasadena, Alhambra, Covina, El Monte, La Puenta  </td>
		</tr><tr>
			<td>628</td><td class="tz">-8</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">PST</td><td>California</td><td class="hidden-xs">Overlay of 415 Area Code  </td>
		</tr><tr>
			<td>629</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Tennessee</td><td class="hidden-xs">Overlay of 615 Area Code  </td>
		</tr><tr>
			<td>630</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Illinois</td><td class="hidden-xs">La Grange, Roselle, Hinsdale, Downers Grove, Naperville, Lombard, Elmhurst, Aurora, Wheaton  </td>
		</tr><tr>
			<td>631</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>New York</td><td class="hidden-xs">Suffolk County: Brentwood, Farmingdale, Central Islip, Riverhead, Huntington, Deer Park, Amityville, Lindenhurst  </td>
		</tr><tr>
			<td>636</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Missouri</td><td class="hidden-xs">St. Charles, Fenton, Harvester, Chesterfield, Manchester  </td>
		</tr><tr>
			<td>641</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Iowa</td><td class="hidden-xs">Fairfield, Mason City, Grinnell, Newton, Knoxville  </td>
		</tr><tr>
			<td>646</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>New York</td><td class="hidden-xs">Overlay of 212 Area Code (New York City).  </td>
		</tr><tr>
			<td>650</td><td class="tz">-8</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">PST</td><td>California</td><td class="hidden-xs">South San Francisco, Palo Alto, San Mateo, Mountain View, Redwood City  </td>
		</tr><tr>
			<td>651</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Minnesota</td><td class="hidden-xs">St. Paul, Redwing, Farmington, Eagan, Lino Lakes, North Branch, Roseville Valley  </td>
		</tr><tr>
			<td>657</td><td class="tz">-8</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">PST</td><td>California</td><td class="hidden-xs">Overlay of 714 area code.  </td>
		</tr><tr>
			<td>660</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Missouri</td><td class="hidden-xs">Warrensburg, Kirksville, Sedalia, Chillicothe, Moberly, Marshall  </td>
		</tr><tr>
			<td>661</td><td class="tz">-8</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">PST</td><td>California</td><td class="hidden-xs">Bakersfield, Santa Clarita, Palmdale, Simi Valley  </td>
		</tr><tr>
			<td>662</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Mississippi</td><td class="hidden-xs">Northern Mississippi: Tupelo, Columbus, Greenwood, Greenville, Oxford  </td>
		</tr><tr>
			<td>667</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Maryland</td><td class="hidden-xs">Overlay of 410, 443 area codes  </td>
		</tr><tr>
			<td>669</td><td class="tz">-8</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">PST</td><td>California</td><td class="hidden-xs">Overlay of 408 area code (San Jose, San Jose, Sunnyvale, etc.)  </td>
		</tr><tr>
			<td>678</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Georgia</td><td class="hidden-xs">Overlay of 404 & 770 area codes.  </td>
		</tr><tr>
			<td>681</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>West Virginia</td><td class="hidden-xs">Overlay of 304 Area Code.  </td>
		</tr><tr>
			<td>682</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Texas</td><td class="hidden-xs">Overlay of 817 area code.  </td>
		</tr><tr>
			<td>700</td><td class="tz">0</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center"></td><td></td><td class="hidden-xs">Used for Value Added Special Services, per individual Carrier.  </td>
		</tr><tr>
			<td>701*</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>North Dakota</td><td class="hidden-xs">North Dakota: All regions *(Area Code crosses into MST time zone) </td>
		</tr><tr>
			<td>702</td><td class="tz">-8</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">PST</td><td>Nevada</td><td class="hidden-xs">Las Vegas & Surrounding Areas.  </td>
		</tr><tr>
			<td>703</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Virginia</td><td class="hidden-xs">Arlington, Alexandria, Fairfax, Falls Church, Quantico, Herndon, Vienna  </td>
		</tr><tr>
			<td>704</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>North Carolina</td><td class="hidden-xs">Charlotte, Asheville, Gastonia, Concord, Statesville, Salisbury, Shelby, Monroe  </td>
		</tr><tr>
			<td>706</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Georgia</td><td class="hidden-xs">Augusta, Columbus, Athens, Rome, Dalton  </td>
		</tr><tr>
			<td>707</td><td class="tz">-8</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">PST</td><td>California</td><td class="hidden-xs">Eureka, Napa, Santa Rosa, Petaluma, Vallejo  </td>
		</tr><tr>
			<td>708</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Illinois</td><td class="hidden-xs">Oak Brook, Calumet City, Harvey (South suburbs of Chicago)  </td>
		</tr><tr>
			<td>710</td><td class="tz">0</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center"></td><td></td><td class="hidden-xs">Federal Government Special Services  </td>
		</tr><tr>
			<td>712</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Iowa</td><td class="hidden-xs">Sioux City, Council Bluffs, Spencer, Cherokee, Denison   </td>
		</tr><tr>
			<td>713</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Texas</td><td class="hidden-xs">Houston  </td>
		</tr><tr>
			<td>714</td><td class="tz">-8</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">PST</td><td>California</td><td class="hidden-xs">Anaheim, Santa Anna  </td>
		</tr><tr>
			<td>715</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Wisconsin</td><td class="hidden-xs">Eau Claire, Wausau, Stevens Point, Superior  </td>
		</tr><tr>
			<td>716</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>New York</td><td class="hidden-xs">Buffalo, Niagara Falls, Tonawanda, Williamsville, Jamestown, Lancaster  </td>
		</tr><tr>
			<td>717</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Pennsylvania</td><td class="hidden-xs">Harrisburg, York, Lancaster  </td>
		</tr><tr>
			<td>718</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>New York</td><td class="hidden-xs">New York City (Bronx, Brooklyn, Queens, Staten Island)  </td>
		</tr><tr>
				<td>719</td><td class="tz">-7</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">MST</td><td>Colorado</td><td class="hidden-xs">Colorado Springs, Pueblo, Lamar  </td>
		</tr><tr>
			<td>720</td><td class="tz">-7</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">MST</td><td>Colorado</td><td class="hidden-xs">Overlay of 303 area code.  </td>
		</tr><tr>
			<td>724</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Pennsylvania</td><td class="hidden-xs">Greensburg, Uniontown, Butler, Washington, New Castle, Indiana  </td>
		</tr><tr>
			<td>725</td><td class="tz">-8</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">PST</td><td>Nevada</td><td class="hidden-xs">Overlay of 702 area code  </td>
		</tr><tr>
			<td>727</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Florida</td><td class="hidden-xs">Clearwater, St. Petersburg, New Port Richey, Tarpon Springs  </td>
		</tr><tr>
			<td>731</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Tennessee</td><td class="hidden-xs">Jackson, Dyersburg, Union City, Paris, Brownsville  </td>
		</tr><tr>
			<td>732</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>New Jersey</td><td class="hidden-xs">New Brunswick, Metuchen, Rahway, Perth Amboy, Toms River, Bound Brook  </td>
		</tr><tr>
			<td>734</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Michigan</td><td class="hidden-xs">Ann Arbor, Livonia, Wayne, Wyandotte, Ypsilanti, Plymouth, Monroe  </td>
		</tr><tr>
			<td>737</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Texas</td><td class="hidden-xs">Overlay of 512 area code  </td>
		</tr><tr>
			<td>740</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Ohio</td><td class="hidden-xs">Portsmouth, Newark, Zanesville, Wheeling, Steubenville  </td>
		</tr><tr>
			<td>743</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>North Carolina</td><td class="hidden-xs">Overlay of 336 area code.  </td>
		</tr><tr>
			<td>747</td><td class="tz">-8</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">PST</td><td>California</td><td class="hidden-xs">Overlay of 818 Area Code.  </td>
		</tr><tr>
			<td>754</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Florida</td><td class="hidden-xs">Overlay of 954 Area Code (Ft. Lauderdale Area).  </td>
		</tr><tr>
			<td>757</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Virginia</td><td class="hidden-xs">Norfolk, Newport News, Williamsburgh  </td>
		</tr><tr>
			<td>760</td><td class="tz">-8</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">PST</td><td>California</td><td class="hidden-xs">Oceanside, Palm Springs, Victorville, Escondido, Vista, Palm Desert  </td>
		</tr><tr>
			<td>762</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Georgia</td><td class="hidden-xs">Overlay of 706 Area Code.  </td>
		</tr><tr>
			<td>763</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Minnesota</td><td class="hidden-xs">Brooklyn Park, Coon Rapids, Maple Grove, Plymouth, Cambridge, Blaine, Anoka  </td>
		</tr><tr>
			<td>765</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Indiana</td><td class="hidden-xs">Lafayette, Muncie, Kokomo, Anderson, Richmond, Marion  </td>
		</tr><tr>
			<td>769</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Mississippi</td><td class="hidden-xs">Overlay of 601 area code (Southern Mississippi).  </td>
		</tr><tr>
			<td>770</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Georgia</td><td class="hidden-xs">Norcross, Chamblee, Smyrna, Tucker, Marietta, Alpharetta, Gainesville  </td>
		</tr><tr>
			<td>772</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Florida</td><td class="hidden-xs">Vero Beach, Port St. Lucie, Stuart, Fort Pierce, Sebastian, Hobe Sound, Jensen Beach, Indiantown  </td>
		</tr><tr>
			<td>773</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Illinois</td><td class="hidden-xs">Chicago (except downtown Area is 312)  </td>
		</tr><tr>
			<td>774</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Massachusetts</td><td class="hidden-xs">Overlay of 508 area code.  </td>
		</tr><tr>
			<td>775*</td><td class="tz">-8</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">PST</td><td>Nevada</td><td class="hidden-xs">Reno, Carson City *(Area Code crosses into MST time zone) </td>
		</tr><tr>
			<td>779</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Illinois</td><td class="hidden-xs">Overlay of 815 area code.  </td>
		</tr><tr>
			<td>781</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Massachusetts</td><td class="hidden-xs">Waltham, Lexington, Burlinton, Dedham, Woburn, Lynn, Malden, Saugus, Reading, Braintree, Wellesley  </td>
		</tr><tr>
			<td>785*</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Kansas</td><td class="hidden-xs">Topeka, Lawrence, Manhattan, Salina, Junction *(Area Code crosses into MST time zone) </td>
		</tr><tr>
			<td>786</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Florida</td><td class="hidden-xs">Overlay of 305 area code (Miami).  </td>
		</tr><tr>
			<td>800</td><td class="tz">0</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center"></td><td></td><td class="hidden-xs">Toll Free  </td>
		</tr><tr>
			<td>801</td><td class="tz">-7</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">MST</td><td>Utah</td><td class="hidden-xs">Utah: Salt Lake City, Provo, Ogden, Orem, American Fork, Spanish Fork, Bountiful, Kaysville, Morgan  </td>
		</tr><tr>
			<td>802</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Vermont</td><td class="hidden-xs">Vermont: All regions  </td>
		</tr><tr>
			<td>803</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>South Carolina</td><td class="hidden-xs">Columbia  </td>
		</tr><tr>
			<td>804</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Virginia</td><td class="hidden-xs">Richmond, Lynchburg, Petersburg,   </td>
		</tr><tr>
			<td>805</td><td class="tz">-8</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">PST</td><td>California</td><td class="hidden-xs">Santa Barbara, Thousand Oaks, San Luis Obispo, Ventura, Oxnard, Simi Valley   </td>
		</tr><tr>
			<td>806</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Texas</td><td class="hidden-xs">Amarillo, Lubbock  </td>
		</tr><tr>
			<td>808</td><td class="tz">-10</td><td class="dst">N</td><td class="time"></td><td class="hidden-xs" align="center">HAST</td><td>Hawaii</td><td class="hidden-xs">Hawaii: All regions  </td>
		</tr><tr>
			<td>809</td><td class="tz">-4</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">AST</td><td></td><td class="hidden-xs">Dominican Republic  </td>
		</tr><tr>
			<td>810</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Michigan</td><td class="hidden-xs">Flint, Port Huron, Lapeer, Brighton, Sandusky  </td>
		</tr><tr>
			<td>812*</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Indiana</td><td class="hidden-xs">Evansville, Bloomington, Terre Haute, New Albany *(Area Code crosses into CST time zone) </td>
		</tr><tr>
			<td>813</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Florida</td><td class="hidden-xs">Tampa  </td>
		</tr><tr>
			<td>814</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Pennsylvania</td><td class="hidden-xs">Altoona, Johnstown, Erie, Punxsutawney  </td>
		</tr><tr>
			<td>815</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Illinois</td><td class="hidden-xs">La Salle, Joliet, Rockford, DeKalb  </td>
		</tr><tr>
			<td>816</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Missouri</td><td class="hidden-xs">Kansas City, Saint Joseph, Independence, Gladstone  </td>
		</tr><tr>
			<td>817</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Texas</td><td class="hidden-xs">Fort Worth, Arlington, Euless, Grapevine  </td>
		</tr><tr>
			<td>818</td><td class="tz">-8</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">PST</td><td>California</td><td class="hidden-xs">Van Nuys, Canoga Park, Burbank, San Fernando, Glendale, N. Hollywood  </td>
		</tr><tr>
			<td>828</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>North Carolina</td><td class="hidden-xs">Asheville, Hickory, Morganton, Hendersonville, Lenoir, Boone, Andrews, Murphy, Marble  </td>
		</tr><tr>
			<td>830</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Texas</td><td class="hidden-xs">New Braunfels, Del Rio, Seguin, Kerrville, Eagle Pass, Fredericksburg  </td>
		</tr><tr>
			<td>831</td><td class="tz">-8</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">PST</td><td>California</td><td class="hidden-xs">Monterey, Santa Cruz, Salinas, Hollister, Aptos, Carmel  </td>
		</tr><tr>
			<td>832</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Texas</td><td class="hidden-xs">Overlay of 713, 281 and 346 area codes (Houston).  </td>
		</tr><tr>
			<td>843</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>South Carolina</td><td class="hidden-xs">Charleston, Florence, Myrtle Beach, Hilton Head  </td>
		</tr><tr>
			<td>844</td><td class="tz">0</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center"></td><td></td><td class="hidden-xs">Toll Free  </td>
		</tr><tr>
			<td>845</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>New York</td><td class="hidden-xs">Poughkeepsie, Spring Valley, Newburgh, Kingston, Nyack, Middletown, Brewster, Pearl River  </td>
		</tr><tr>
			<td>847</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Illinois</td><td class="hidden-xs">Northbrook, Skokie, Evanston, Glenview, Waukegan, Desplaines, Elk Grove (North suburbs of Chicago)  </td>
		</tr><tr>
			<td>848</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>New Jersey</td><td class="hidden-xs">Overlay of 732 area code.  </td>
		</tr><tr>
			<td>850*</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Florida</td><td class="hidden-xs">Tallahassee, Pensacola, Fort Walton Beach, Panama City *(Area Code crosses into EST time zone) </td>
		</tr><tr>
			<td>854</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>South Carolina</td><td class="hidden-xs">Overlay of the 843 area code.  </td>
		</tr><tr>
			<td>855</td><td class="tz">0</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center"></td><td></td><td class="hidden-xs">Toll Free  </td>
		</tr><tr>
			<td>856</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>New Jersey</td><td class="hidden-xs">Camden, Haddonfield, Moorestown, Merchantville, Vineland, Laurel Springs  </td>
		</tr><tr>
			<td>857</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Massachusetts</td><td class="hidden-xs">Overlay of 617 area code.  </td>
		</tr><tr>
			<td>858</td><td class="tz">-8</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">PST</td><td>California</td><td class="hidden-xs">La Jolla, Rancho Bernardo, del Mar, Poway, Rancho Penasquitos, Rancho Santa Fe  </td>
		</tr><tr>
			<td>859</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Kentucky</td><td class="hidden-xs">Lexington, Covington, Boone, Winchester, Richmond, Danville, Mount Sterling  </td>
		</tr><tr>
			<td>860</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Connecticut</td><td class="hidden-xs">Hartford, New London, Norwich, Middletown, Bristol  </td>
		</tr><tr>
			<td>862</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>New Jersey</td><td class="hidden-xs">Overlay of 973 area code.  </td>
		</tr><tr>
			<td>863</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Florida</td><td class="hidden-xs">Lakeland, Winter Haven, Lake Wales, Sebring, Haines City, Bartow, Avon Park, Okeechobee, Wachula  </td>
		</tr><tr>
			<td>864</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>South Carolina</td><td class="hidden-xs">Greenville, Spartanburg, Anderson  </td>
		</tr><tr>
			<td>865</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Tennessee</td><td class="hidden-xs">Knoxville, Maryville, Oak Ridge, Sevierville, Gatlinburg, Concord, Powell  </td>
		</tr><tr>
			<td>866</td><td class="tz">0</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center"></td><td></td><td class="hidden-xs">Toll Free  </td>
		</tr><tr>
			<td>870</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Arkansas</td><td class="hidden-xs">Texarkana, Jonesboro, Pine Bluff, El Dorado  </td>
		</tr><tr>
			<td>872</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Illinois</td><td class="hidden-xs">Overlay of 312 and 773 area codes.  </td>
		</tr><tr>
			<td>877</td><td class="tz">0</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center"></td><td></td><td class="hidden-xs">Toll Free  </td>
		</tr><tr>
			<td>878</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Pennsylvania</td><td class="hidden-xs">Overlay of 412, 724 area codes.  </td>
		</tr><tr>
			<td>880</td><td class="tz">0</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center"></td><td></td><td class="hidden-xs">Carribean to U.S. or Canada OR Canada to U.S. Toll Free  </td>
		</tr><tr>
			<td>881</td><td class="tz">0</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center"></td><td></td><td class="hidden-xs">Carribean to U.S. or Canada OR Canada to U.S. Toll Free  </td>
		</tr><tr>
			<td>888</td><td class="tz">0</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center"></td><td></td><td class="hidden-xs">Toll Free  </td>
		</tr><tr>
			<td>900</td><td class="tz">0</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center"></td><td></td><td class="hidden-xs">Used for Mass Calling Value Added Information  </td>
		</tr><tr>
			<td>901</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Tennessee</td><td class="hidden-xs">Memphis, Collierville  </td>
		</tr><tr>
			<td>903</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Texas</td><td class="hidden-xs">Longview, Tyler, Texarkana, Paris, Kilgore, Sherman, Denison  </td>
		</tr><tr>
			<td>904</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Florida</td><td class="hidden-xs">Jacksonville, Saint Augustine, Orange Park, Fernandina Beach  </td>
		</tr><tr>
			<td>906*</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Michigan</td><td class="hidden-xs">Marquette, Iron Mountain, Houghton, Sault Ste Marie *(Area Code crosses into CST time zone) </td>
		</tr><tr>
			<td>907*</td><td class="tz">-9</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">AKST</td><td>Alaska</td><td class="hidden-xs">Alaska: All regions *(Area Code crosses into HAST time zone) </td>
		</tr><tr>
			<td>908</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>New Jersey</td><td class="hidden-xs">Elizabeth, New Brunswick, Somerville, Freehold, Unionville, Plainfield  </td>
		</tr><tr>
			<td>909</td><td class="tz">-8</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">PST</td><td>California</td><td class="hidden-xs">San Bernardino, Ontario, Upland, Pomona, Riverside, Colton, Chino  </td>
		</tr><tr>
			<td>910</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>North Carolina</td><td class="hidden-xs">Fayetteville, Wilmington, Jacksonville, Lumberton, Laurinburg, Southern Pines  </td>
		</tr><tr>
			<td>912</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Georgia</td><td class="hidden-xs">Savannah, Macon, Waycross, Brunswick, Statesboro, Vidalia  </td>
		</tr><tr>
			<td>913</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Kansas</td><td class="hidden-xs">Melrose, Kansas City  </td>
		</tr><tr>
			<td>914</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>New York</td><td class="hidden-xs">Westchester, Monroe, Mount Vernon, Mount Kisco, Pleasantville  </td>
		</tr><tr>
			<td>915*</td><td class="tz">-7</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">MST</td><td>Texas</td><td class="hidden-xs">El Paso, Socorro, Fabens, Dell City, Van Horn, Fort Bliss *(Area Code crosses into CST time zone) </td>
		</tr><tr>
			<td>916</td><td class="tz">-8</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">PST</td><td>California</td><td class="hidden-xs">Sacramento, Roseville, Fair Oaks, Folsom, Elk Grove, South Placer  </td>
		</tr><tr>
			<td>917</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>New York</td><td class="hidden-xs">Overlay of 212/718 (New York City - cellular/pager only)  </td>
		</tr><tr>
			<td>918</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Oklahoma</td><td class="hidden-xs">Tulsa, Broken Arrow, Muskogee, Bartlesville, McAlester  </td>
		</tr><tr>
			<td>919</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>North Carolina</td><td class="hidden-xs">Raleigh, Durham, Cary, Chapel Hill, Goldsboro, Apex, Sanford, Wake Forest  </td>
		</tr><tr>
			<td>920</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Wisconsin</td><td class="hidden-xs">Green Bay, Appleton, Racine, Fond du Lac, Oshkosh, Sheboygan  </td>
		</tr><tr>
			<td>925</td><td class="tz">-8</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">PST</td><td>California</td><td class="hidden-xs">Walnut Creek, Pleasanton, Concord, Livermore, Bishop Ranch, Danville, Antioch  </td>
		</tr><tr>
			<td>928*</td><td class="tz">-7</td><td class="dst">N</td><td class="time"></td><td class="hidden-xs" align="center">MST</td><td>Arizona</td><td class="hidden-xs">Yuma, Flagstaff, Prescott, Sedona, Bullhead City, Kingman, Lake Havasu City *(Area Code crosses into MST time zone) </td>
		</tr><tr>
			<td>929</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>New York</td><td class="hidden-xs">Overlay of 718 Area Code (Bronx, Brooklyn, Queens, Staten Island).  </td>
		</tr><tr>
			<td>930*</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Indiana</td><td class="hidden-xs">Overlay of 812 Area Code *(Area Code crosses into CST time zone) </td>
		</tr><tr>
			<td>931</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Tennessee</td><td class="hidden-xs">Clarksville, Shelbyville, Cookeville, Columbia  </td>
		</tr><tr>
			<td>934</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>New York</td><td class="hidden-xs">Overlay of 631 Area Code.  </td>
		</tr><tr>
			<td>936</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Texas</td><td class="hidden-xs">Conroe, Nacogdoches, Huntsville, Lufkin, Madisonville  </td>
		</tr><tr>
			<td>937</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Ohio</td><td class="hidden-xs">Dayton, Springfield, Bellefontaine, Beavercreek, Franklin  </td>
		</tr><tr>
			<td>938</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Alabama</td><td class="hidden-xs">Overlay of 256 area code.  </td>
		</tr><tr>
			<td>940</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Texas</td><td class="hidden-xs">Wichita Falls, Denton, Gainesville  </td>
		</tr><tr>
			<td>941</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Florida</td><td class="hidden-xs">Sarasota, Bradenton, Venice, Port Charlotte, Punta Gorda  </td>
		</tr><tr>
			<td>947</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Michigan</td><td class="hidden-xs">Overlay of 248 Area Code.  </td>
		</tr><tr>
			<td>949</td><td class="tz">-8</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">PST</td><td>California</td><td class="hidden-xs">Irvine, Saddleback Valley, Newport Beach, Capistrano Valley, Laguna Beach  </td>
		</tr><tr>
			<td>951</td><td class="tz">-8</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">PST</td><td>California</td><td class="hidden-xs">Riverside, Corona, Temecula, Arlington, Hemet, Moreno Valley, Murietta, Sun City, Elsinore  </td>
		</tr><tr>
			<td>952</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Minnesota</td><td class="hidden-xs">Bloomington, Burnsville, Eden Prairie, Minnetonka, Edina, St. Louis Park, Apple Valley  </td>
		</tr><tr>
			<td>954</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Florida</td><td class="hidden-xs">Fort Lauderdale, Hollywood, Pompano Beach, Deerfield Beach, Coral Springs  </td>
		</tr><tr>
			<td>956</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Texas</td><td class="hidden-xs">Brownsville, Laredo, McAllen, Harlingen, Edinburg    </td>
		</tr><tr>
			<td>959</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Connecticut</td><td class="hidden-xs">Overlay of 860 Area Code  </td>
		</tr><tr>
			<td>970</td><td class="tz">-7</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">MST</td><td>Colorado</td><td class="hidden-xs">Fort Collins, Greeley, Grand Junction, Loveland, Durango, Vail  </td>
		</tr><tr>
			<td>971</td><td class="tz">-8</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">PST</td><td>Oregon</td><td class="hidden-xs">Overlay of 503 area code.  </td>
		</tr><tr>
			<td>972</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Texas</td><td class="hidden-xs">Grand Prairie, Addison, Irving, Richardson, Plano, Carrollton  </td>
		</tr><tr>
			<td>973</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>New Jersey</td><td class="hidden-xs">Newark, Morristown, Paterson, Passaic, Orange, Bloomfield, Caldwell, Whippany  </td>
		</tr><tr>
			<td>978</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Massachusetts</td><td class="hidden-xs">Lowell, Lawrence, Billerica, Concord, Wilmington, Sudbury, Fitchburg, Peabody, Andover, Beverly, Danvers  </td>
		</tr><tr>
			<td>979</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Texas</td><td class="hidden-xs">Bryan, Lake Jackson, Freeport, Brenham, Bay City, El Campo  </td>
		</tr><tr>
			<td>980</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>North Carolina</td><td class="hidden-xs">Overlay of 704 Area Code.  </td>
		</tr><tr>
			<td>984</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>North Carolina</td><td class="hidden-xs">Overlay of 919 area code.  </td>
		</tr><tr>
			<td>985</td><td class="tz">-6</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">CST</td><td>Louisiana</td><td class="hidden-xs">Houma, Slidell, Hammond, Morgan City, Mandeville, Covington, Thibodaux, Bogalusa, St. Charles  </td>
		</tr><tr>
			<td>989</td><td class="tz">-5</td><td class="dst">Y</td><td class="time"></td><td class="hidden-xs" align="center">EST</td><td>Michigan</td><td class="hidden-xs">Saginaw, Bay City, Midland, Mount Pleasant  </td>
		</tr>
	</table>
</div>

	    </div>

        <br />
        <footer class="jumbotron" style="padding:0;">
            <div style="background-color:#d2d2d2; padding-top:10px;">
                <div class="container"><p>&copy; 1994-2017  GreatData.com</p></div>
            </div>
            <div class="container">
                <div class="row">
                    <div class="col-md-10">
                        <p><a href="content/resourcefiles/privacy.txt" id="ctl00_A3" target="_blank">Privacy Policy</a></p>
                        <p><a href="ContactUs.aspx">Contact</a> </p>
                        <p><a href="pdf/LicenseAgreement-Simplified.pdf" id="ctl00_A2" target="_blank">Software/Data License</a> </p>
                        <p>ZIP Code Database</p>
                    </div>
                    <div class="col-md-2">
                        <div class="socialmedia"><a class="fb" href="https://www.facebook.com/pages/GreatDatacom/313618730969" ></a><br /><a class="in" href="http://www.linkedin.com/company/greatdata.com" ></a><br /><a class="tw" href="http://twitter.com/#!/greatdata" ></a></div>
                    </div>
                </div>
            </div>
            <!-- &nbsp; - &nbsp; <span id="ctl00_lblPageName"></span> -->
       </footer>

        <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.2/js/bootstrap.min.js"></script>

    <script type="text/javascript">
        $(document).ready(function () {

            //calc time 1 sec after load, then 2 sec, then 3 sec... then end. The reason for the delay is the func sometimes beats the grid data its relying on when loading.
            AddTimeToPage();
            setTimeout(AddTimeToPage, 500);
            setTimeout(AddTimeToPage, 1000);
            setTimeout(AddTimeToPage, 3000);

            //check every 5 sec: if min change, re calc time - this simply refreses the times
            var tmpMin = new Date().getMinutes();
            setInterval(function () {
                var updMin = new Date().getMinutes()
                if (tmpMin != updMin) {
                    AddTimeToPage();
                    tmpMin = updMin
                }
            }, 4000);

        });


        function AddTimeToPage() {
            $("#ctl00_cph_body_gvTZAC tr").each(function () {
                var tzcell = $(this).closest('tr').find('.tz');
                var timecell = $(this).closest('tr').find('.time');
                var dstcell = $(this).closest('tr').find('.dst');
                timecell.html(calcTime(tzcell.html(), dstcell.html() ));
            });
        }


        function calcTime(offset, dstflag) {
            var d = new Date(); // create Date object for current location
            var utc;

            if (dstflag == 'N') { //does not observe dst (exa. much of arizona)
                utc = d.getTime() + (d.getTimezoneOffset() * 60000);
            }
            else {
                utc = d.getTime() + (AdjustForDST(d.getTimezoneOffset()) * 60000); // 1) convert to msec 2)add local time zone offset 3) adjust for Daylight Savings Time 4)get UTC time in msec
            }

            var nd = new Date(utc + (3600000 * offset)); // create new Date object for different city using supplied offset

            var hours = ((nd.getHours() + 11) % 12 + 1); //12hr clock
            var min = (nd.getMinutes() < 10 ? '0' : '') + nd.getMinutes() //also adds leading zero if min < 10

            var ampm = nd.getHours() > 11 ? "PM" : "AM";
            return hours + ":" + min + " " + ampm;  // return newdate/time as a string
        }

        function AdjustForDST(offset) {
            //see http://forums.asp.net/t/1963152.aspx?Javascript+local+time+with+daylight+saving+considered+in+
            var d = new Date();
            var jan = new Date(d.getFullYear(), 0, 1);
            var jul = new Date(d.getFullYear(), 6, 1);

            if (offset < Math.max(jan.getTimezoneOffset(), jul.getTimezoneOffset())) {
                return Math.max(jan.getTimezoneOffset(), jul.getTimezoneOffset());
            }
            else {
                return offset;
            }
        }




    </script>




        <script>
            //prevent closing signin box when user clicks mouse inside usr/pwd textboxes
            $('#ctl00_mnuLogin_UserName').click(function (e) {
                e.stopPropagation();
            });

            $('#ctl00_mnuLogin_Password').click(function (e) {
                e.stopPropagation();
            });
        </script>
        <script>
            (function (i, s, o, g, r, a, m) {
                i['GoogleAnalyticsObject'] = r; i[r] = i[r] || function () {
                    (i[r].q = i[r].q || []).push(arguments)
                }, i[r].l = 1 * new Date(); a = s.createElement(o), m = s.getElementsByTagName(o)[0]; a.async = 1; a.src = g; m.parentNode.insertBefore(a, m)
            })(window, document, 'script', '//www.google-analytics.com/analytics.js', 'ga');

            ga('create', 'UA-4212606-2', 'auto');
            ga('send', 'pageview');
        </script>

        <script>
            var quotes = [
                " “It’s always been a pleasure working with GreatData.” – American Express",
                " “Your professionalism and attention to detail should be the benchmark for all companies...” – Kinkos",
                " “I find much value in your data ...” – General Electric ",
                " “Your company has always been very responsive to our inquiries for service and help. ” – Owens Corning",
                " “I have been using your data well over 12 years now. I love it ...” – GE Capital",
                " “Very accurate data.” – Ford&nbsp;",
                " “Your data has proved invaluable for us many times.” – FedEx",
                " “I appreciate the consistently excellent service and data!” – Nestlé Purina" ];

            //quotes.sort(function () { return 0.5 - Math.random(); }); //this emulates a random sort, but not very efficient or very random
            quotes = shuffle(quotes);

            $.each(quotes, function (index, quote) {
                var element = $('<div class="item">' + quote + '&nbsp</div>')
                if (index === 0) {
                    element.addClass("active");
                }
                element.appendTo($('.carousel-inner'));
            });

            $('#quoteslide').carousel({
                interval: 6000

            })

            function shuffle(array) {
                var currentIndex = array.length, temporaryValue, randomIndex;

                // While there remain elements to shuffle...
                while (0 !== currentIndex) {

                    // Pick a remaining element...
                    randomIndex = Math.floor(Math.random() * currentIndex);
                    currentIndex -= 1;

                    // And swap it with the current element.
                    temporaryValue = array[currentIndex];
                    array[currentIndex] = array[randomIndex];
                    array[randomIndex] = temporaryValue;
                }

                return array;
            }

        </script>



    </form>


</body>
</html>`
