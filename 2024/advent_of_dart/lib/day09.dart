mixin Day09 {
  int day09a(List<String> lines) {
    int result = 0;

    final disk = lines[0];
    int front = 0;
    int fid = 0;
    int pos = 0;
    int back = disk.length - 1;
    int backPtr = 0;
    int backSize = int.parse(disk[back]);
    while (front < back) {
      // calc checksum of next file
      for (int i = 0; i < int.parse(disk[front]); i++) {
        print('adding front file: $pos * $fid');
        result += fid * pos++;
      }
      front++; // move to next block on disk (free space)
      // calc checksum of free space (read from end of file backwards)
      fid = back ~/ 2;
      for (int i = 0; i < int.parse(disk[front]); i++) {
        print('adding rear file: $pos * $fid');
        result += fid * pos++;
        if (++backPtr >= backSize) {
          // out of back
          back -= 2;
          fid = back ~/ 2;
          backPtr = 0;
          backSize = int.parse(disk[back]);
        }
      }

      front++; // move front to next block
      fid = front ~/ 2;
    }

    // finish writing the last rear file
    while (backPtr < backSize) {
      print('adding rear file: $pos * $fid (end)');
      result += fid * pos++;
      backPtr++;
    }

    return result;
  }

  int day09b(List<String> lines) {
    int result = 0;

    final disk = lines[0].split('');
    List<(int, int)> files = [];
    for (int i = 0; i < disk.length; i++) {
      final idx = i.isEven ? i ~/ 2 : -1;
      files.add((idx, int.parse(disk[i])));
    }

    for (int i = disk.length-1; i >= 0; i -= 2) {
      final idx = i ~/ 2;
      final size = int.parse(disk[i]);
      for (int j = 0; j < files.length; j++) {
        if (files[j].$1 == idx) break;
        if (files[j].$1 != -1) continue;

        final freeSize = files[j].$2;
        if (freeSize >= size) {
          // found a slot. change existing file to empty
          files[files.indexOf((idx, size))] = (-1, size);
          if (freeSize == size) { // perfect fit, just overwrite j
            files[j] = (idx, size);
          } else { // excess free space. split into two entries
            files[j] = (-1, freeSize - size);
            files.insert(j, (idx, size));
          }
          break;
        }
      }
    }

    // calc checksum
    int idx = 0;
    for (final file in files) {
      final fileIdx = file.$1;
      if (fileIdx == -1) { // free space
        idx += file.$2;
        continue;
      }
      for (int j = 0; j < file.$2; j++) {
        result += idx++ * fileIdx;
      }
    }

    return result;
  }
}
