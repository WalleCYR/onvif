package replay

import (
	"github.com/WalleCYR/onvif/xsd"
	"github.com/WalleCYR/onvif/xsd/onvif"
)

// Capabilities type
type Capabilities struct {

	// Indicator that the Device supports reverse playback as defined in the ONVIF Streaming Specification.

	ReversePlayback bool `xml:"ReversePlayback,attr,omitempty"`

	// The list contains two elements defining the minimum and maximum valid values supported as session timeout in seconds.

	SessionTimeoutRange xsd.FloatAttrList `xml:"SessionTimeoutRange,attr,omitempty"`

	// Indicates support for RTP/RTSP/TCP.

	RTP_RTSP_TCP bool `xml:"RTP_RTSP_TCP,attr,omitempty"`

	// If playback streaming over WebSocket is supported, this shall return the RTSP WebSocket URI as described in Streaming Specification Section 5.1.1.5.

	RTSPWebSocketUri xsd.AnyURI `xml:"RTSPWebSocketUri,attr,omitempty"`
}

// GetServiceCapabilities type
type GetServiceCapabilities struct {
	XMLName string `xml:"GetServiceCapabilities"`
}

// GetServiceCapabilitiesResponse type
type GetServiceCapabilitiesResponse struct {
	XMLName string `xml:"GetServiceCapabilitiesResponse"`

	// The capabilities for the replay service is returned in the Capabilities element.
	Capabilities Capabilities `xml:"Capabilities,omitempty"`
}

// GetReplayUri type
type GetReplayUri struct {
	XMLName onvif.Name `xml:"GetReplayUri"`

	// Specifies the connection parameters to be used for the stream. The URI that is returned may depend on these parameters.
	StreamSetup onvif.StreamSetup `xml:"StreamSetup,omitempty"`

	// The identifier of the recording to be streamed.
	RecordingToken onvif.ReferenceToken `xml:"RecordingToken,omitempty"`
}

// GetReplayUriResponse type
type GetReplayUriResponse struct {
	XMLName onvif.Name `xml:"GetReplayUriResponse"`

	// The URI to which the client should connect in order to stream the recording.
	Uri xsd.AnyURI `xml:"Uri,omitempty"`
}

// SetReplayConfiguration type
type SetReplayConfiguration struct {
	XMLName onvif.Name `xml:"SetReplayConfiguration"`

	// Description of the new replay configuration parameters.
	Configuration ReplayConfiguration `xml:"Configuration,omitempty"`
}

// SetReplayConfigurationResponse type
type SetReplayConfigurationResponse struct {
	XMLName onvif.Name `xml:"SetReplayConfigurationResponse"`
}

// GetReplayConfiguration type
type GetReplayConfiguration struct {
	XMLName onvif.Name `xml:"GetReplayConfiguration"`
}

// GetReplayConfigurationResponse type
type GetReplayConfigurationResponse struct {
	XMLName onvif.Name `xml:"GetReplayConfigurationResponse"`

	// The current replay configuration parameters.
	Configuration ReplayConfiguration `xml:"Configuration,omitempty"`
}

// ReplayConfiguration type
type ReplayConfiguration struct {

	// The RTSP session timeout.
	SessionTimeout xsd.Duration `xml:"SessionTimeout,omitempty"`
}
