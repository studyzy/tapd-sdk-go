package tapd

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/studyzy/tapd-sdk-go/model"
)

// GetImage 获取图片下载链接
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/attachment/get_image.html
func (c *Client) GetImage(ctx context.Context, req *model.GetImageRequest) (*model.ImageInfo, error) {
	data, err := c.doGet(ctx, "/files/get_image", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseOne[model.ImageInfo](data, "Attachment")
}

// GetAttachments 获取附件列表（含下载链接）
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/attachment/get_attachments.html
func (c *Client) GetAttachments(ctx context.Context, req *model.GetAttachmentsRequest) ([]model.Attachment, error) {
	data, err := c.doGet(ctx, "/attachments", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseList[model.Attachment](data, "Attachment")
}

// GetOneAttachment 获取单个附件下载链接
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/attachment/get_one_attachment.html
func (c *Client) GetOneAttachment(ctx context.Context, req *model.GetOneAttachmentRequest) (*model.Attachment, error) {
	data, err := c.doGet(ctx, "/files/get_one_attachment", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseOne[model.Attachment](data, "Attachment")
}

// DownloadDocument 获取单个文档下载链接
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/attachment/download_document.html
func (c *Client) DownloadDocument(ctx context.Context, req *model.DownloadDocumentRequest) (json.RawMessage, error) {
	data, err := c.doGet(ctx, "/documents/down", req.ToParams())
	if err != nil {
		return nil, err
	}
	return data, nil
}

// UploadAttachment 上传附件
// API 文档：https://open.tapd.cn/document/api-doc/API文档/mini_api_reference/attachment/upload_attachment.html
func (c *Client) UploadAttachment(ctx context.Context, req *model.UploadAttachmentRequest, filename string, fileReader io.Reader) (*model.Attachment, error) {
	pr, pw := io.Pipe()
	writer := multipart.NewWriter(pw)

	go func() {
		defer pw.Close()
		defer writer.Close()

		writer.WriteField("workspace_id", req.WorkspaceID)
		writer.WriteField("type", req.Type)
		writer.WriteField("custom_field", req.CustomField)
		writer.WriteField("entry_id", req.EntryID)
		if req.Owner != "" {
			writer.WriteField("owner", req.Owner)
		}

		part, err := writer.CreateFormFile("file", filename)
		if err != nil {
			pw.CloseWithError(err)
			return
		}
		if _, err := io.Copy(part, fileReader); err != nil {
			pw.CloseWithError(err)
			return
		}
	}()

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/files/upload_attachment", pr)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	c.applyAuth(httpReq)
	httpReq.Header.Set("Content-Type", writer.FormDataContentType())

	data, err := c.doRequest(httpReq)
	if err != nil {
		return nil, err
	}

	return parseOne[model.Attachment](data, "Attachment")
}

// UploadImageBase64 上传 base64 图片
// API 文档：https://open.tapd.cn/document/api-doc/API文档/mini_api_reference/attachment/upload_image_base64.html
func (c *Client) UploadImageBase64(ctx context.Context, req *model.UploadImageBase64Request) (*model.Attachment, error) {
	data, err := c.doPost(ctx, "/files/upload_image_base64", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseOne[model.Attachment](data, "Attachment")
}

// DownloadAttachment 获取单个附件下载链接
// API 文档：https://open.tapd.cn/document/api-doc/API文档/mini_api_reference/attachment/get_one_attachment.html
func (c *Client) DownloadAttachment(ctx context.Context, req *model.DownloadAttachmentRequest) (*model.Attachment, error) {
	data, err := c.doGet(ctx, "/attachments/down", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseOne[model.Attachment](data, "Attachment")
}
