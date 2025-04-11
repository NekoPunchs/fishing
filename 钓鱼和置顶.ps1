$proc = Start-Process fishing.exe -PassThru
Start-Sleep -Seconds 5
$handle = (Get-Process -Id $proc.Id).MainWindowHandle
$signature = @'
[DllImport("user32.dll")]
public static extern bool SetWindowPos(
    IntPtr hWnd,
    IntPtr hWndInsertAfter,
    int X,
    int Y,
    int cx,
    int cy,
    uint uFlags);
'@
$type = Add-Type -MemberDefinition $signature -Name SetWindowPosition -Namespace SetWindowPos -Using System.Text -PassThru
$alwaysOnTop = New-Object -TypeName System.IntPtr -ArgumentList (-1)
$type::SetWindowPos($handle, $alwaysOnTop, 0, 0, 0, 0, 0x0003)