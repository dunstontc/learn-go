## Web Programming Synonymous Terms
- router
- request router
- multiplexer
- mux
- servemux
- server
- http router
- http request router
- http multiplexer
- http mux
- http servemux
- http server

***

## Request & response

Request and response messages are similar. Both messages consist of:

- a request/response line
- zero or more header lines
- a blank line (ie, CRLF by itself)
- an optional message body

***

### HTTP request

Request
- request line
- headers
- optional message body

Request-Line
- Method SP Request-URI SP HTTP-Version CRLF

example request line:
- GET /path/to/file/index.html HTTP/1.1

***

### HTTP response

#### Reponse
- status line
- headers
- optional message body

#### Status-Line
- `HTTP-Version SP Status-Code SP Reason-Phrase CRLF`

#### example status line:
```http
HTTP/1.1 200 OK
```

***

### Headers
[List of header fields](https://en.wikipedia.org/wiki/List_of_HTTP_header_fields)

***

### Inspect
- you can use google chrome dev tools / network
- you can use CURL at the command line

```bash
curl -v golang.org
```

***

## Documentation

### [Hypertext Transfer Protocol](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol)

Development of HTTP was initiated by Tim Berners-Lee at CERN in 1989. Standards development of HTTP was coordinated by the Internet Engineering Task Force (IETF) and the World Wide Web Consortium (W3C), culminating in the publication of a series of Requests for Comments (RFCs). 

The first definition of HTTP/1.1, the version of HTTP in common use, occurred in RFC 2068 in 1997, although this was obsoleted by RFC 2616 in 1999. In 2014, RFC2616 was replaced by multiple RFCs (7230-7237)

A later version, the successor HTTP/2, was standardized in 2015, and is now supported by major web servers.

### [Request For Comment (RFC)](https://en.wikipedia.org/wiki/Request_for_Comments)

A Request for Comments (RFC) is a type of publication from the Internet Engineering Task Force (IETF) and the Internet Society (ISOC), the principal technical development and standards-setting bodies for the Internet.

An RFC is authored by engineers and computer scientists in the form of a memorandum describing methods, behaviors, research, or innovations applicable to the working of the Internet and Internet-connected systems. It is submitted either for peer review or simply to convey new concepts, information, or (occasionally) engineering humor. The IETF adopts some of the proposals published as RFCs as Internet Standards.

Request for Comments documents were invented by Steve Crocker in 1969 to help record unofficial notes on the development of ARPANET. RFCs have since become official documents of Internet specifications, communications protocols, procedures, and events.

### [Internet Engineering Task Force (IETF)](https://en.wikipedia.org/wiki/Internet_Engineering_Task_Force)

The Internet Engineering Task Force (IETF) develops and promotes voluntary Internet standards, in particular the standards that comprise the Internet protocol suite (TCP/IP). It is an open standards organization, with no formal membership or membership requirements. All participants and managers are volunteers, though their work is usually funded by their employers or sponsors.

The IETF started out as an activity supported by the U.S. federal government, but since 1993 it has operated as a standards development function under the auspices of the Internet Society, an international membership-based non-profit organization.

### [List of RFCs](https://en.wikipedia.org/wiki/List_of_RFCs)


***

## RFC 7230

HTTP was created for the World Wide Web (WWW) architecture and has
evolved over time to support the scalability needs of a worldwide
hypertext system.  Much of that architecture is reflected in the
terminology and syntax productions used to define HTTP.

### 2.1.  Client/Server Messaging

HTTP is a stateless request/response protocol that operates by
exchanging messages across a reliable TRANSPORT- or
SESSION-layer "CONNECTION".


An HTTP "CLIENT" is a program that establishes a CONNECTION to a server
for the purpose of sending one or more HTTP requests.

An HTTP "SERVER" is a program that accepts CONNECTIONS in order to service HTTP requests by sending HTTP responses.

The terms "client" and "server" refer only to the roles that these
programs perform for a particular connection.  The same program might
act as a client on some connections and a server on others.  The term
"user agent" refers to any of the various client programs that
initiate a request, including (but not limited to) browsers, spiders
(web-based robots), command-line tools, custom applications, and
mobile apps.  The term "origin server" refers to the program that can
originate authoritative responses for a given target resource.  The
terms "sender" and "recipient" refer to any implementation that sends
or receives a given message, respectively.

HTTP relies upon the Uniform Resource Identifier (URI) standard
to indicate the target resource and relationships between resources.
Messages are passed in a format similar to that used by Internet mail
and the Multipurpose Internet Mail Extensions (MIME) (see Appendix A of
[RFC7231] for the differences between HTTP and MIME messages).

Most HTTP communication consists of a retrieval request (GET) for a
representation of some resource identified by a URI.  In the simplest
case, this might be accomplished via a single bidirectional
CONNECTION (===) between the user agent (UA) and the origin
server (O).

```
	request   >
UA ======================================= O
						   <   response
```
A client sends an HTTP REQUEST to a server in the form of a REQUEST
MESSAGE, beginning with a REQUEST-LINE that includes a method, URI,
and protocol version, followed by header fields
containing request modifiers, client information, and representation
metadata, an empty line to indicate the end of the
header section, and finally a message body containing the payload
body (if any, Section 3.3).

A server responds to a client's request by sending one or more HTTP
RESPONSE MESSAGES, each beginning with a STATUS LINE that includes
the protocol version, a success or error code, and textual reason
phrase, possibly followed by header fields containing
server information, resource metadata, and representation metadata,
an empty line to indicate the end of the header
section, and finally a message body containing the payload body (if
any, Section 3.3).

A connection might be used for multiple REQUEST/RESPONSE exchanges,
as defined in Section 6.3.

The following example illustrates a typical message exchange for a
GET request on the URI
"http://www.example.com/hello.txt":

### Client request:

```http
GET /hello.txt HTTP/1.1

User-Agent: curl/7.16.3 libcurl/7.16.3 OpenSSL/0.9.7l zlib/1.2.3

Host: www.example.com

Accept-Language: en, mi
```

### Server response:

```http
HTTP/1.1 200 OK

Date: Mon, 27 Jul 2009 12:28:53 GMT

Server: Apache

Last-Modified: Wed, 22 Jul 2009 19:15:56 GMT

ETag: "34aa387-d-1568eb00"

Accept-Ranges: bytes

Content-Length: 51

Vary: Accept-Encoding

Content-Type: text/plain

Hello World! My payload includes a trailing CRLF.
```

### 2.2.  Implementation Diversity

When considering the design of HTTP, it is easy to fall into a trap
of thinking that all user agents are general-purpose browsers and all
origin servers are large public websites.  That is not the case in
practice.  Common HTTP USER AGENTS include household appliances,
stereos, scales, firmware update scripts, command-line programs,
mobile apps, and communication devices in a multitude of shapes and
sizes.  Likewise, common HTTP ORIGIN SERVERS include home automation
units, configurable networking components, office machines,
autonomous robots, news feeds, traffic cameras, ad selectors, and
video-delivery platforms.

The term "USER AGENT" does not imply that there is a human user
directly interacting with the software agent at the time of a
request.  In many cases, a user agent is installed or configured to
run in the background and save its results for later inspection (or
save only a subset of those results that might be interesting or
erroneous).  Spiders, for example, are typically given a start URI
and configured to follow certain behavior while crawling the Web as a
hypertext graph.


## OSI

<table class="wikitable" style="margin: 1em auto 1em auto;">
  <tbody>
    <tr>
      <th colspan="5">OSI Model</th>
    </tr>
    <tr>
      <th colspan="2">Layer</th>
      <th><a href="/wiki/Protocol_data_unit" title="Protocol data unit">Protocol data unit</a> (PDU)
      </th>
      <th style="width:30em;">Function<sup id="cite_ref-3" class="reference"><a href="#cite_note-3">[3]</a></sup></th>
    </tr>
    <tr>
      <th rowspan="4">Host<br>layers</th>
      <td style="background:#d8ec9b;">7.&nbsp;<a href="/wiki/Application_layer" title="Application layer">Application</a></td>
      <td style="background:#d8ec9c;" rowspan="3"><a href="/wiki/Data_(computing)" title="Data (computing)">Data</a></td>
      <td style="background:#d8ec9c;"><small>High-level <a href="/wiki/API" class="mw-redirect" title="API">APIs</a>, including resource sharing, remote file access</small></td>
    </tr>
    <tr>
      <td style="background:#d8ec9b;">6.&nbsp;<a href="/wiki/Presentation_layer" title="Presentation layer">Presentation</a></td>
      <td style="background:#d8ec9b;"><small>Translation of data between a networking service and an application; including <a href="/wiki/Character_encoding" title="Character encoding">character encoding</a>, <a href="/wiki/Data_compression" title="Data compression">data compression</a> and <a href="/wiki/Encryption" title="Encryption">encryption/decryption</a></small></td>
    </tr>
    <tr>
      <td style="background:#d8ec9b;">5. <a href="/wiki/Session_layer" title="Session layer">Session</a></td>
      <td style="background:#d8ec9b;"><small>Managing communication <a href="/wiki/Session_(computer_science)" title="Session (computer science)">sessions</a>, i.e. continuous exchange of information in the form of multiple back-and-forth transmissions between two nodes</small></td>
    </tr>
    <tr>
      <td style="background:#e7ed9c;">4. <a href="/wiki/Transport_layer" title="Transport layer">Transport</a></td>
      <td style="background:#e7ed9c;"><a href="/wiki/Packet_segmentation" title="Packet segmentation">Segment</a>, <a href="/wiki/Datagram" title="Datagram">Datagram</a></td>
      <td style="background:#e7ed9c;"><small>Reliable transmission of data segments between points on a network, including <a href="/wiki/Packet_segmentation" title="Packet segmentation">segmentation</a>, <a href="/wiki/Acknowledgement_(data_networks)" title="Acknowledgement (data networks)">acknowledgement</a> and <a href="/wiki/Multiplexing" title="Multiplexing">multiplexing</a></small></td>
    </tr>
    <tr>
      <th rowspan="3">Media<br>layers</th>
      <td style="background:#eddc9c;">3. <a href="/wiki/Network_layer" title="Network layer">Network</a></td>
      <td style="background:#eddc9c;"><a href="/wiki/Network_packet" title="Network packet">Packet</a></td>
      <td style="background:#eddc9c;"><small>Structuring and managing a multi-node network, including <a href="/wiki/Address_space" title="Address space">addressing</a>, <a href="/wiki/Routing" title="Routing">routing</a> and <a href="/wiki/Network_traffic_control" title="Network traffic control">traffic control</a></small></td>
    </tr>
    <tr>
      <td style="background:#e9c189;">2. <a href="/wiki/Data_link_layer" title="Data link layer">Data link</a></td>
      <td style="background:#e9c189;"><a href="/wiki/Frame_(networking)" title="Frame (networking)">Frame</a></td>
      <td style="background:#e9c189;"><small>Reliable transmission of data frames between two nodes connected by a physical layer</small></td>
    </tr>
    <tr>
      <td style="background:#e9988a;">1. <a href="/wiki/Physical_layer" title="Physical layer">Physical</a></td>
      <td style="background:#e9988a;"><a href="/wiki/Symbol_rate#Symbols" title="Symbol rate">Symbol</a></td>
      <td style="background:#e9988a;"><small>Transmission and reception of raw bit streams over a physical medium</small></td>
    </tr>
  </tbody>
</table>
