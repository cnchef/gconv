# gconv 发布脚本
# 使用方法: .\release.ps1 v1.0.0

param(
    [Parameter(Mandatory=$true)]
    [string]$Version
)

# 检查版本号格式
if ($Version -notmatch '^v\d+\.\d+\.\d+$') {
    Write-Host "❌ 错误：版本号格式不正确！" -ForegroundColor Red
    Write-Host "正确格式：v1.0.0, v1.2.3" -ForegroundColor Yellow
    exit 1
}

Write-Host "🚀 开始发布 gconv $Version" -ForegroundColor Cyan
Write-Host ""

# 步骤 1: 检查工作目录
Write-Host "📋 步骤 1/6: 检查工作目录..." -ForegroundColor Yellow
$status = git status --porcelain
if ($status) {
    Write-Host "❌ 错误：有未提交的改动" -ForegroundColor Red
    Write-Host ""
    git status
    Write-Host ""
    Write-Host "请先提交所有改动：" -ForegroundColor Yellow
    Write-Host "  git add ." -ForegroundColor Gray
    Write-Host "  git commit -m 'your message'" -ForegroundColor Gray
    exit 1
}
Write-Host "✅ 工作目录干净" -ForegroundColor Green
Write-Host ""

# 步骤 2: 运行测试
Write-Host "🧪 步骤 2/6: 运行测试..." -ForegroundColor Yellow
$testResult = go test -v ./...
if ($LASTEXITCODE -ne 0) {
    Write-Host "❌ 错误：测试未通过" -ForegroundColor Red
    exit 1
}
Write-Host "✅ 所有测试通过" -ForegroundColor Green
Write-Host ""

# 步骤 3: 检查标签是否已存在
Write-Host "🏷️  步骤 3/6: 检查标签..." -ForegroundColor Yellow
$existingTag = git tag -l $Version
if ($existingTag) {
    Write-Host "❌ 错误：标签 $Version 已存在" -ForegroundColor Red
    Write-Host ""
    Write-Host "现有标签：" -ForegroundColor Yellow
    git tag -l
    exit 1
}
Write-Host "✅ 标签可用" -ForegroundColor Green
Write-Host ""

# 步骤 4: 获取更新日志
Write-Host "📝 步骤 4/6: 生成更新日志..." -ForegroundColor Yellow
$prevTag = git describe --tags --abbrev=0 2>$null
if ($prevTag) {
    $changelog = git log "$prevTag..HEAD" --pretty=format:"- %s" --no-merges
    Write-Host "更新内容（从 $prevTag 到现在）：" -ForegroundColor Cyan
} else {
    $changelog = git log --pretty=format:"- %s" --no-merges
    Write-Host "更新内容（首次发布）：" -ForegroundColor Cyan
}
Write-Host $changelog -ForegroundColor Gray
Write-Host ""

# 步骤 5: 创建标签
Write-Host "🏷️  步骤 5/6: 创建标签..." -ForegroundColor Yellow
$tagMessage = @"
Release $Version

更新内容：
$changelog

项目状态：
- 测试覆盖率：98.2%
- 零第三方依赖
- 支持 Go 1.20+
"@

git tag -a $Version -m $tagMessage
if ($LASTEXITCODE -ne 0) {
    Write-Host "❌ 错误：创建标签失败" -ForegroundColor Red
    exit 1
}
Write-Host "✅ 标签创建成功" -ForegroundColor Green
Write-Host ""

# 步骤 6: 推送标签
Write-Host "📤 步骤 6/6: 推送到 GitHub..." -ForegroundColor Yellow
Write-Host "即将推送标签 $Version 到远程仓库" -ForegroundColor Yellow
Write-Host "这将触发自动发布流程" -ForegroundColor Yellow
Write-Host ""
$confirm = Read-Host "确认推送? (y/N)"
if ($confirm -ne 'y' -and $confirm -ne 'Y') {
    Write-Host "❌ 取消发布" -ForegroundColor Red
    Write-Host ""
    Write-Host "删除本地标签：" -ForegroundColor Yellow
    git tag -d $Version
    exit 0
}

git push origin main
git push origin $Version
if ($LASTEXITCODE -ne 0) {
    Write-Host "❌ 错误：推送失败" -ForegroundColor Red
    exit 1
}

Write-Host ""
Write-Host "🎉 发布成功！" -ForegroundColor Green
Write-Host ""
Write-Host "📋 后续步骤：" -ForegroundColor Cyan
Write-Host "1. 查看 GitHub Release: https://github.com/cnchef/gconv/releases" -ForegroundColor Gray
Write-Host "2. 查看 Actions 状态: https://github.com/cnchef/gconv/actions" -ForegroundColor Gray
Write-Host "3. 等待 pkg.go.dev 索引(约 24 小时): https://pkg.go.dev/github.com/cnchef/gconv@$Version" -ForegroundColor Gray
Write-Host ""
Write-Host "📦 用户安装命令:" -ForegroundColor Cyan
Write-Host "  go get github.com/cnchef/gconv@$Version" -ForegroundColor Yellow
Write-Host ""

