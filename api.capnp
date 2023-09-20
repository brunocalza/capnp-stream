using Go = import "/go.capnp";

@0x9cf9878fd3dd8473;

$Go.package("stream");
$Go.import("stream");

interface FileUploader {

	upload @0 (filename :Text) -> (callback :Callback);
	interface Callback {
		write @0 (chunk :Data);
		done @1 ();
	}
}