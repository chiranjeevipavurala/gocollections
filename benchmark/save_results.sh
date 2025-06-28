#!/bin/bash

# Go Collections Benchmark Results Saver
# This script runs benchmarks and saves results to a file with metadata

# Configuration
RESULTS_DIR="benchmark_results"
TIMESTAMP=$(date +"%Y-%m-%d_%H-%M-%S")
SYSTEM_INFO=$(uname -s)_$(uname -m)
GO_VERSION=$(go version | awk '{print $3}')
RESULTS_FILE="${RESULTS_DIR}/benchmark_results_${SYSTEM_INFO}_${TIMESTAMP}.txt"

# Create results directory if it doesn't exist
mkdir -p "$RESULTS_DIR"

echo "🚀 Go Collections Benchmark Suite - Results Saver"
echo "=================================================="
echo ""

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Error: Go is not installed or not in PATH"
    exit 1
fi

# Check if we're in the right directory
if [ ! -f "go.mod" ]; then
    echo "❌ Error: Please run this script from the project root directory"
    exit 1
fi

echo "🔧 System Information:"
echo "   Go version: $GO_VERSION"
echo "   Architecture: $(go env GOARCH)"
echo "   OS: $(go env GOOS)"
echo "   CPU cores: $(nproc 2>/dev/null || sysctl -n hw.ncpu 2>/dev/null || echo "unknown")"
echo "   Results file: $RESULTS_FILE"
echo ""

# Create header with system information
cat > "$RESULTS_FILE" << EOF
# Go Collections Benchmark Results
# Generated: $(date)
# System: $(uname -a)
# Go version: $GO_VERSION
# Architecture: $(go env GOARCH)
# OS: $(go env GOOS)
# CPU cores: $(nproc 2>/dev/null || sysctl -n hw.ncpu 2>/dev/null || echo "unknown")

EOF

echo "📊 Running benchmarks and saving results..."
echo "   This may take a few minutes..."

# Run benchmarks and append to file
{
    echo "## All Benchmarks"
    echo "\`\`\`"
    go test -bench=. -benchtime=2s ./benchmark/
    echo "\`\`\`"
    echo ""
} >> "$RESULTS_FILE"

# Run memory benchmarks
{
    echo "## Memory Allocation Benchmarks"
    echo "\`\`\`"
    go test -bench=. -benchmem -benchtime=2s ./benchmark/
    echo "\`\`\`"
    echo ""
} >> "$RESULTS_FILE"

# Generate summary
{
    echo "## Performance Summary"
    echo ""
    echo "### Fastest Operations (< 20 ns/op)"
    echo "Based on the benchmark results above, these are the fastest operations:"
    echo ""
    echo "### Good Operations (20-50 ns/op)"
    echo "These operations perform well:"
    echo ""
    echo "### Slower Operations (> 50 ns/op)"
    echo "These operations are slower but may be expected for their complexity:"
    echo ""
    echo "### Recommendations"
    echo "- Use ArrayList for random access operations"
    echo "- Use LinkedList for frequent insertions/deletions at ends"
    echo "- Use HashMap for maximum performance without ordering"
    echo "- Use LinkedHashMap when you need insertion order (minimal overhead)"
    echo "- Use TreeMap when you need sorted keys"
    echo ""
} >> "$RESULTS_FILE"

echo "✅ Benchmark results saved to: $RESULTS_FILE"
echo ""

# Create a latest results symlink
LATEST_LINK="${RESULTS_DIR}/latest_results.txt"
rm -f "$LATEST_LINK"
ln -s "$(basename "$RESULTS_FILE")" "$LATEST_LINK"

echo "🔗 Latest results available at: $LATEST_LINK"
echo ""

# Show file size
FILE_SIZE=$(du -h "$RESULTS_FILE" | cut -f1)
echo "📁 Results file size: $FILE_SIZE"
echo ""

# Optional: Show a quick summary
echo "📈 Quick Summary (first few lines):"
echo "----------------------------------------"
head -20 "$RESULTS_FILE"
echo "----------------------------------------"
echo ""

echo "🎯 To view all results:"
echo "   cat $RESULTS_FILE"
echo ""
echo "🎯 To compare with previous results:"
echo "   diff $RESULTS_DIR/benchmark_results_*.txt"
echo ""
echo "🎯 To view latest results:"
echo "   cat $LATEST_LINK" 