import os
import ctypes

# ファイルを完全に削除する関数
def secure_delete(filepath):
    # ファイルのサイズを取得
    file_size = os.path.getsize(filepath)

    # ファイルの中身をゼロで上書き（ゼロフィリング）
    with open(filepath, 'wb') as file:
        file.write(b'0' * file_size)

    # ファイルのメタデータを削除
    ctypes.windll.kernel32.SetFileAttributesW(filepath, 2)

    # ファイルを物理的に上書きし、ディスク上から削除（シュレッディング）
    ctypes.windll.kernel32.SetFileInformationByHandle(
        ctypes.windll.kernel32.CreateFileW(filepath, 0x80000000, 0, None, 3, 0x80, None),
        2, None, 0
    )

# 使用例
file_path = r'C:\path\to\file.txt'  # 削除したいファイルのパス
secure_delete(file_path)
