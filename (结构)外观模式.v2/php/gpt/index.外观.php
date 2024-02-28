<?php
namespace php\gpt\facade;

class VideoFile {
    private $filename;

    public function __construct($filename) {
        $this->filename = $filename;
    }

    public function save() {
        // 保存视频文件的代码...
    }
}

class OggCompressionCodec {
    public function compress($video) {
        // 这里是实际的压缩算法...
    }
}

class MPEG4CompressionCodec {
    public function compress($video) {
        // 这里是另一种压缩算法...
    }
}

class CodecFactory {
    public static function getCodec($format) {
        if ($format === 'mp4') {
            return new MPEG4CompressionCodec();
        } else {
            return new OggCompressionCodec();
        }
    }
}

class VideoConversionFacade {
    public function convertVideo($filename, $format) {
        $file = new VideoFile($filename);
        $sourceCodec = CodecFactory::getCodec($format);
        // 可能会有更多的逻辑，选择正确的编解码器，创建位流等
        $destinationCodec = ($format === 'mp4') ? new OggCompressionCodec() : new MPEG4CompressionCodec();

        $sourceCodec->compress($file);

        $file->save(); // 假设这里是保存压缩过后的文件

        return $file;
    }
}

$facade = new VideoConversionFacade();
$convertedVideo = $facade->convertVideo("example_video.mkv", "mp4");

/**
在这个例子中，VideoConversionFacade 类就是一个外观。它隐藏了视频压缩和文件处理的复杂性，提供了一个简单的方法 convertVideo 来完成整个流程。

外观模式通常在以下情况下使用：

当你要为一个复杂的子系统提供一个简单的接口时。
当你想要解耦系统和使用它的客户端代码时。
当你需要一个层次化的结构去表示子系统时。
使用外观模式有助于降低客户端与子系统之间的耦合度，从而使系统的扩展和维护变得更简单。
 */
